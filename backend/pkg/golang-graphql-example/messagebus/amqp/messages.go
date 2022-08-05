package amqpbusmessage

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofrs/uuid"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/middlewares"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	"github.com/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
)

func (as *amqpService) Publish(
	ctx context.Context,
	message *amqp091.Publishing,
	publishCfg *PublishConfigInput,
) (err error) {
	// Increase active request counter
	as.signalHandlerSvc.IncreaseActiveRequestCounter()
	// Get trace from context
	trace := as.tracingSvc.StartChildTraceOrTraceFromContext(ctx, tracingPublishOperation)
	// Defer the closing trace
	defer func() {
		// Decrease active request counter
		as.signalHandlerSvc.DecreaseActiveRequestCounter()
		// Check if error is set
		if err != nil {
			// Mark trace as in error
			trace.MarkAsError()
			// Increase failed counter
			as.metricsSvc.IncreaseFailedAMQPPublishedMessage(publishCfg.Exchange, publishCfg.RoutingKey)
		} else {
			// Increase success counter
			as.metricsSvc.IncreaseSuccessfullyAMQPPublishedMessage(publishCfg.Exchange, publishCfg.RoutingKey)
		}
		// Close trace
		trace.Finish()
	}()

	// Add correlation id in message if not set
	if message.CorrelationId == "" {
		// Get request id
		reqID := middlewares.GetRequestIDFromContext(ctx)
		// Check if request id is set
		if reqID != "" {
			// Use it
			message.CorrelationId = reqID
		} else {
			// Generate new id
			id, err2 := generateCorrelationID()
			// Check error
			if err2 != nil {
				return errors.WithStack(err2)
			}

			// Save it
			message.CorrelationId = id
		}
	}

	// Add info to trace
	trace.SetTags(map[string]interface{}{
		"exchange":       publishCfg.Exchange,
		"routing-key":    publishCfg.RoutingKey,
		"correlation-id": message.CorrelationId,
		"message-id":     message.MessageId,
		"priority":       message.Priority,
		"type":           message.Type,
	})

	// Create logger fields
	fields := map[string]interface{}{
		"exchange":       publishCfg.Exchange,
		"routing-key":    publishCfg.RoutingKey,
		"correlation-id": message.CorrelationId,
		"message-id":     message.MessageId,
		"priority":       message.Priority,
		"type":           message.Type,
	}
	// Check if trace id exists
	if trace.GetTraceID() != "" {
		// Set it in log
		fields[log.LogTraceIDField] = trace.GetTraceID()
	}
	// Get logger from context
	logger := log.GetLoggerFromContext(ctx)
	// Add fields in logger
	logger = logger.WithFields(fields)

	// Check if headers are set, otherwise create it
	if message.Headers == nil {
		message.Headers = amqp091.Table{}
	}
	// Create headers
	err = as.injectTracedHeaders(trace, message.Headers)
	// Check error
	if err != nil {
		return err
	}

	// Initialize retry send delay
	sendDelayDur := defaultRetryDelay
	// Check if params have it set
	if publishCfg.RetryDelay != 0 {
		sendDelayDur = publishCfg.RetryDelay
	}

	// Initialize timeout duration
	timeoutDuration := defaultPublishTimeout
	// Check if params have it set
	if publishCfg.Timeout != 0 {
		timeoutDuration = publishCfg.Timeout
	}
	// Create new context for timeout
	// This is done to avoid closing main context
	timeoutCtx, cancelF := context.WithTimeout(context.TODO(), timeoutDuration)
	// Defer cancel if it completes before timeout elapses
	defer cancelF()

	// Loop
	for {
		// Create a channel to save errors
		errChan := make(chan error, 1)

		// Create publish result channel
		pubResChan := make(chan bool, 1)

		// Check if channel isn't opened or present
		if as.publisherChannel == nil || as.publisherChannel.IsClosed() {
			errChan <- errors.New("publisher channel not present or closed")
		} else {
			// Publish
			confirmation, pErr := as.publisherChannel.PublishWithDeferredConfirm(
				publishCfg.Exchange,
				publishCfg.RoutingKey,
				publishCfg.Mandatory,
				publishCfg.Immediate,
				*message,
			)
			// Check error
			if pErr != nil {
				// Check if channel is closed, if yes, put it in retry
				if as.publisherChannel.IsClosed() {
					errChan <- pErr
				} else {
					// In this case, return error
					// As check that channel is opened before, this case have a lower rank
					// Error here must happened when configuration is incorrect or something else in broker
					return errors.WithStack(pErr)
				}
			} else {
				// Start a routine for wait response
				go func() {
					d := confirmation.Wait()
					pubResChan <- d
				}()
			}
		}

		select {
		// Timeout case
		case <-timeoutCtx.Done():
			return ErrPublishTimeoutReached
		// Error management
		case err := <-errChan:
			logger.Error(errors.Wrap(err, "error detected when tried to publish, retrying after delay"))
			time.Sleep(sendDelayDur)
		// Published
		case ack := <-pubResChan:
			// Check if ack is ok
			if ack {
				// Finish
				logger.Debug("message successfully published")

				return nil
			}

			logger.Debug("message published but not ack, retrying after delay")
		// Retry
		case <-time.After(sendDelayDur):
			logger.Debug("publish retry delay reached, retrying")
		}
	}
}

func (as *amqpService) Consume(
	ctx context.Context,
	getConsumeCfg func() *ConsumeConfigInput,
	cb func(ctx context.Context, delivery *amqp091.Delivery) error,
) error {
	// Loop
	for {
		// Get configuration
		// This is function to allow the support of hot reloading the configuration.
		consumeCfg := getConsumeCfg()

		// Initialize retry send delay
		sendDelayDur := defaultRetryDelay
		// Check if params have it set
		if consumeCfg.RetryDelay != 0 {
			sendDelayDur = consumeCfg.RetryDelay
		}

		// Create consumer logger
		logger := log.GetLoggerFromContext(ctx).WithFields(map[string]interface{}{
			"queue": consumeCfg.QueueName,
		})

		// Check if system isn't closing
		if as.signalHandlerSvc.IsStoppingSystem() {
			// Closing in progress
			// Just stop consume
			logger.Info("consume stopped, system is stopping")

			return nil
		}

		// Check if channel isn't opened or present
		if as.consumerChannel == nil || as.consumerChannel.IsClosed() {
			// Create error
			err := errors.New("error detected when tried to consumer: consumer channel not present or closed, retrying after delay")
			// Log
			logger.Error(err)
			// Wait
			time.Sleep(sendDelayDur)

			continue
		}

		// Get hostname for consumer tag
		hostname, err := os.Hostname()
		// Check error
		if err != nil {
			return errors.WithStack(err)
		}
		// Build consumer tag
		consumerTag := fmt.Sprintf("%s-%s", consumeCfg.ConsumerPrefix, hostname)

		// Append to consumer tags list if not present
		as.appendToConsumerTags(consumerTag)

		// Consume
		deliveries, cErr := as.consumerChannel.Consume(
			consumeCfg.QueueName,
			consumerTag,
			consumeCfg.AutoAck,
			consumeCfg.Exclusive,
			consumeCfg.NoLocal,
			consumeCfg.NoWait,
			consumeCfg.Args,
		)
		// Check error
		if cErr != nil {
			// Check if channel is closed, if yes, put it in retry
			if as.consumerChannel.IsClosed() {
				// Create error
				err := errors.New("error detected when tried to consumer: consumer channel not present or closed, retrying after delay")
				// Log
				logger.Error(err)
				// Wait
				time.Sleep(sendDelayDur)

				continue
			}

			// In this case, return error
			// As check that channel is opened before, this case have a lower rank
			// Error here must happened when configuration is incorrect or something else in broker
			return errors.WithStack(cErr)
		}

		logger.Debug("Waiting for consumer message")

		// Loop over deliveries
		for d := range deliveries {
			// Increase active request counter
			as.signalHandlerSvc.IncreaseActiveRequestCounter()

			// Create handler
			handler := func() (err error) {
				// Extract trace from message
				trace, err := as.extractTraceFromHeaders(d.Headers)
				// Check error
				if err != nil {
					return err
				}
				// Defer to close trace
				defer func() {
					// Check error
					if err != nil {
						trace.MarkAsError()
					}

					trace.Finish()
				}()

				// Check if correlation id is set
				// Otherwise, create it
				if d.CorrelationId == "" {
					// Generate new id
					id, err2 := generateCorrelationID()
					// Check error
					if err2 != nil {
						return errors.WithStack(err2)
					}

					// Save it
					d.CorrelationId = id
				}

				// Set tags in trace
				trace.SetTags(map[string]interface{}{
					"queue":          consumeCfg.QueueName,
					"consumer-tag":   d.ConsumerTag,
					"routing-key":    d.RoutingKey,
					"correlation-id": d.CorrelationId,
					"message-id":     d.MessageId,
					"priority":       d.Priority,
					"type":           d.Type,
					"redelivered":    d.Redelivered,
				})

				// Create fields
				fields := map[string]interface{}{
					"correlation_id": d.CorrelationId,
					"consumer_tag":   d.ConsumerTag,
					"routing_key":    d.RoutingKey,
					"message_id":     d.MessageId,
					"priority":       d.Priority,
					"type":           d.Type,
					"redelivered":    d.Redelivered,
				}
				// Check if trace id exists
				if trace.GetTraceID() != "" {
					// Set it in log
					fields[log.LogTraceIDField] = trace.GetTraceID()
				}
				// Update
				logger = logger.WithFields(fields)
				// Create new context with logger
				cbCtx := log.SetLoggerToContext(ctx, logger)
				// Set trace in context
				cbCtx = tracing.SetTraceToContext(cbCtx, trace)

				// Log
				logger.Debug("start consuming message")

				// Call handler
				err = as.consumeDeliveryHandler(cbCtx, trace, &d, cb)
				// Check error
				if err != nil {
					logger.Error(errors.Wrap(err, "message consumed failed with error"))

					// Calculate Requeue option
					// Initialize
					requeue := true
					// Check if option is set
					if consumeCfg.RequeueOnNackFn != nil {
						requeue = consumeCfg.RequeueOnNackFn(&d, err)
					}

					// Nack message
					err = d.Nack(false, requeue)
					// Check error
					// This may arrive when worker is disconnected
					if err != nil {
						logger.Error(errors.Wrap(err, "cannot nack consumed message"))
						// Stop
						return nil
					}

					// Increase failed counter
					as.metricsSvc.IncreaseFailedAMQPConsumedMessage(
						consumeCfg.QueueName,
						d.ConsumerTag,
						d.RoutingKey,
					)
				} else {
					// Ack message
					err = d.Ack(false)
					// Check error
					// This may arrive when worker is disconnected
					if err != nil {
						logger.Error(errors.Wrap(err, "cannot ack consumed message"))
						// Stop
						return nil
					}

					logger.Info("message successfully consumed")
					// Increase success counter
					as.metricsSvc.IncreaseSuccessfullyAMQPConsumedMessage(
						consumeCfg.QueueName,
						d.ConsumerTag,
						d.RoutingKey,
					)
				}

				// Default
				return nil
			}

			// Call handler
			err := handler()
			// Check error
			if err != nil {
				logger.Error(err)
			}

			// Decrease active request counter
			as.signalHandlerSvc.DecreaseActiveRequestCounter()
		}
	}
}

func (as *amqpService) consumeDeliveryHandler(
	ctx context.Context,
	trace tracing.Trace,
	d *amqp091.Delivery,
	cb func(ctx context.Context, delivery *amqp091.Delivery) error,
) (err error) {
	// Defer close and status
	defer func() {
		// Try to recover if any
		if err2 := recover(); err2 != nil {
			err = errors.New(fmt.Sprintf("%v", err2))
		}
		// Check if error is set
		if err != nil {
			// Mark trace as in error
			trace.MarkAsError()
		}
		// Close trace
		trace.Finish()
	}()

	// Call the callback
	err = cb(ctx, d)
	// Check error
	if err != nil {
		return err
	}

	// Default
	return nil
}

func (as *amqpService) extractTraceFromHeaders(h amqp091.Table) (tracing.Trace, error) {
	// Create map string string
	headers := map[string]string{}

	// Loop over input headers
	for k, v := range h {
		// Check if value is a string
		switch v := v.(type) { //nolint: gocritic // Ignore because can't do this in if
		case string:
			headers[k] = v
		}
	}

	// Extract
	return as.tracingSvc.ExtractFromTextMapAndStartSpan(headers, tracingConsumeOperation)
}

func (as *amqpService) injectTracedHeaders(trace tracing.Trace, headers amqp091.Table) error {
	// Create headers
	h := map[string]string{}
	// Use inject in headers
	err := trace.InjectInTextMap(h)
	// Check error
	if err != nil {
		return err
	}

	// Create amqp headers
	for k, v := range h {
		headers[k] = v
	}

	// Return
	return nil
}

func generateCorrelationID() (string, error) {
	// Get uuid
	uid, err := uuid.NewV4()
	// Check error
	if err != nil {
		return "", errors.WithStack(err)
	}

	return uid.String(), nil
}
