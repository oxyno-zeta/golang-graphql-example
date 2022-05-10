package amqpbusmessage

import (
	"os"
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/signalhandler"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	"github.com/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/thoas/go-funk"
)

type amqpService struct {
	logger              log.Logger
	cfgManager          config.Manager
	tracingSvc          tracing.Service
	signalHandlerSvc    signalhandler.Client
	metricsSvc          metrics.Client
	publisherConnection *amqp.Connection
	publisherChannel    *amqp.Channel
	consumerConnection  *amqp.Connection
	consumerChannel     *amqp.Channel
	consumerTags        []string
}

func (as *amqpService) Reconnect() error {
	// Close connection and channels
	err := as.Close()
	// Check error
	if err != nil {
		return err
	}

	// Reconnect
	return as.Connect()
}

func (as *amqpService) Close() error {
	// Check if consumer channel is opened
	if !as.consumerChannel.IsClosed() {
		// Closing channel
		err := as.consumerChannel.Close()
		// Check error
		if err != nil && !errors.Is(err, amqp091.ErrClosed) {
			return errors.WithStack(err)
		}
	}

	// Check if publish channel is opened
	if !as.publisherChannel.IsClosed() {
		// Just closing publisher channel as no consumer are in
		err := as.publisherChannel.Close()
		// Check error
		if err != nil && !errors.Is(err, amqp091.ErrClosed) {
			return errors.WithStack(err)
		}
	}

	// Check if publisher connection is opened
	if !as.publisherConnection.IsClosed() {
		// Closing publisher connection
		err := as.publisherConnection.Close()
		// Check error
		if err != nil && !errors.Is(err, amqp091.ErrClosed) {
			return errors.WithStack(err)
		}
	}

	// Check if consumer connection is opened
	if !as.consumerConnection.IsClosed() {
		// Closing consumer connection
		err := as.consumerConnection.Close()
		// Check error
		if err != nil && !errors.Is(err, amqp091.ErrClosed) {
			return errors.WithStack(err)
		}
	}

	// Default
	return nil
}

func (as *amqpService) Connect() error {
	as.logger.Debugf("Trying to connect AMQP bus")

	// Connect publisher
	err := as.connectPublisher()
	// Check error
	if err != nil {
		return err
	}

	// Connect consumer
	err = as.connectConsumer()
	// Check error
	if err != nil {
		return err
	}

	// Start reconnect as routine
	go as.reconnect(func() *amqp.Connection { return as.publisherConnection }, as.connectPublisher)
	go as.reconnect(func() *amqp.Connection { return as.consumerConnection }, as.connectConsumer)

	as.logger.Info("Successfully connected to AMQP broker")

	// Default
	return nil
}

// Reconnect only the connection.
// Channels cannot die if not closed by application itself.
// Don't manage this case in this client.
func (as *amqpService) reconnect(getConnection func() *amqp.Connection, connect func() error) {
	// Infinite loop
	for {
		// Listen for closed events
		errReason := <-getConnection().NotifyClose(make(chan *amqp.Error))

		// Check if reason is set, if not set, it is because the close is coming from application side
		if errReason == nil {
			as.logger.Debug("Reconnection handler have detected an application closing connection event => skip reconnection")

			break
		}

		as.logger.Error(errors.Wrap(errReason, "Attempting to reconnect to AMQP broker because connection was closed due to error"))

		// Loop for reconnect
		for {
			// Wait a bit
			time.Sleep(defaultReconnectWaitingDuration)
			// Call connect
			err := connect()
			// Check if error is empty, meaning that connection is done
			if err == nil {
				// Reconnect is done
				as.logger.Info("Reconnection to AMQP broker successful")

				break
			}

			as.logger.Error(errors.Wrap(err, "Reconnection to AMQP broker failed"))
		}
	}
}

func (as *amqpService) connectPublisher() error {
	// Call internal connect
	conn, chann, err := as.connect()
	// Check error
	if err != nil {
		return err
	}

	// Configure channel to be in confirm publish messages mode
	// This will allow the notify channel to work
	err = chann.Confirm(false)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Save result
	as.publisherConnection = conn
	as.publisherChannel = chann

	return nil
}

func (as *amqpService) connectConsumer() error {
	// Call internal connect
	conn, chann, err := as.connect()
	// Check error
	if err != nil {
		return err
	}

	// Save result
	as.consumerConnection = conn
	as.consumerChannel = chann

	return nil
}

func (as *amqpService) connect() (*amqp.Connection, *amqp.Channel, error) {
	// Get configuration
	cfg := as.cfgManager.GetConfig()

	// Get AMQP message bus config
	amqpCfg := cfg.AMQP

	// Initialize heartbeat
	var heartbeat time.Duration
	// Check if heartbeat is configured
	if amqpCfg.Connection.HeartbeatDuration != "" {
		// Parse
		heartbeatP, err := time.ParseDuration(amqpCfg.Connection.HeartbeatDuration)
		// Check error
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}

		// Save
		heartbeat = heartbeatP
	}

	// Initialize extra args
	extraArgs := amqpCfg.Connection.ExtraArgs
	// Check if it is set
	if extraArgs == nil {
		extraArgs = map[string]interface{}{}
	}
	// Check if connection_name exists in the connection extra args
	// If it doesn't, override to add hostname as connection_name
	if _, ok := extraArgs["connection_name"]; !ok {
		// Get hostname
		hostname, err := os.Hostname()
		// Check error
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}

		extraArgs["connection_name"] = hostname
	}

	// Create AMQP connection configuration
	connACfg := amqp.Config{
		FrameSize:  amqpCfg.Connection.FrameSize,
		ChannelMax: amqpCfg.Connection.ChannelMax,
		Heartbeat:  heartbeat,
		Properties: extraArgs,
	}
	// Check if heartbeat duration is set
	if amqpCfg.Connection.HeartbeatDuration != "" {
		// Try to parse
		dur, err := time.ParseDuration(amqpCfg.Connection.HeartbeatDuration)
		// Check error
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}

		// Save
		connACfg.Heartbeat = dur
	}

	as.logger.Debugf("Trying to establish connection to AMQP bus")
	// Connect
	conn, err := amqp.DialConfig(amqpCfg.Connection.URL.Value, connACfg)
	// Check error
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	as.logger.Debugf("Trying to create a channel on connection to AMQP bus")
	// Create channel with configuration
	chann, err := as.createConfiguredChannel(conn, amqpCfg.ChannelQos)
	// Check error
	if err != nil {
		return nil, nil, err
	}

	// Apply setup
	err = as.setup(amqpCfg, chann)
	// Check error
	if err != nil {
		return nil, nil, err
	}

	return conn, chann, nil
}

func (as *amqpService) createConfiguredChannel(conn *amqp.Connection, channelQosCfg *config.AMQPChannelQosConfig) (*amqp.Channel, error) {
	// Create channel
	chann, err := conn.Channel()
	// Check error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Apply Qos if possible
	// Check if Qos configuration is set
	if channelQosCfg != nil {
		// Apply configuration
		err = chann.Qos(
			channelQosCfg.PrefetchCount,
			channelQosCfg.PrefetchSize,
			channelQosCfg.Global,
		)
		// Check error
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	// Default
	return chann, nil
}

func (as *amqpService) CancelAllConsumers() error {
	// Loop over all consumer tags
	for _, ct := range as.consumerTags {
		// Cancel consumer
		// Note: There isn't any error when there is a cancel on a tag that
		// doesn't exists.
		err := as.consumerChannel.Cancel(ct, false)
		// Check error
		if err != nil {
			return err
		}
	}

	// Default
	return nil
}

func (as *amqpService) Ping() error {
	// Check if in progress or disconnected management in progress
	if as.publisherConnection == nil || as.consumerConnection == nil {
		return errors.New("connection to AMQP broker is closed or not initialized")
	}

	// Get publisher connection status
	isPublisherClosed := as.publisherConnection.IsClosed()
	// Get consumer connection status
	isConsumerClosed := as.consumerConnection.IsClosed()
	// Check status
	if isPublisherClosed || isConsumerClosed {
		return errors.New("connection to AMQP broker is closed")
	}

	// Default case
	return nil
}

func (as *amqpService) appendToConsumerTags(newConsumerTag string) {
	// Add it only if array isn't containing data
	if !funk.ContainsString(as.consumerTags, newConsumerTag) {
		as.consumerTags = append(as.consumerTags, newConsumerTag)
	}
}
