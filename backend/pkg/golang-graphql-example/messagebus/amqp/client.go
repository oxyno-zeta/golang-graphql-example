package amqpbusmessage

import (
	"context"
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/signalhandler"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	"github.com/rabbitmq/amqp091-go"
)

var defaultReconnectWaitingDuration = 200 * time.Millisecond
var defaultPublishTimeout = 10 * time.Second
var defaultRetryDelay = time.Second
var tracingPublishOperation = "amqp:publish"
var tracingConsumeOperation = "amqp:consume"

// ErrPublishTimeoutReached is the error thrown when the publish timeout is over.
var ErrPublishTimeoutReached = errors.NewInternalServerError("timeout reached")

// ErrMessageNotJSON is the error thrown when the message utils parse function is called on a non "application/json" message.
var ErrMessageNotJSON = errors.NewInvalidInputError("input haven't the json content type")

// PublishConfigInput represents the publish configuration input.
type PublishConfigInput struct {
	// Exchange is the exchange name where the message is published.
	Exchange string
	// RoutingKey is the published message routing key.
	RoutingKey string
	// Mandatory (see mandatory in AMQP)
	Mandatory bool
	// Immediate (see immediate in AMQP)
	Immediate bool
	// Timeout represents the timeout to wait for publishing a message and have an ack from server including retries.
	// If max timeout is reach, the ErrPublishTimeoutReached is raised.
	// If not set, a default timeout is set to 10 seconds.
	Timeout time.Duration
	// RetryDelay represents the wait delay to consider that a message haven't been sent to server.
	// This can be considered as the timeout between a message is sent and no ack have been detected from server.
	// This will trigger a message publish retry.
	// If not set, a default delay is set to 1 second.
	// Don't go below this limit as a message can takes time to be ack.
	RetryDelay time.Duration
}

// ConsumeConfigInput represents the consume configuration input.
type ConsumeConfigInput struct {
	// QueueName is the queue name for consume.
	QueueName string
	// ConsumerPrefix is the prefix used for the consumer tag in AMQP consumer.
	// The suffix is the hostname.
	ConsumerPrefix string
	// AutoAck
	AutoAck bool
	// Exclusive
	Exclusive bool
	// NoLocal
	NoLocal bool
	// NoWait
	NoWait bool
	// Args
	Args amqp091.Table
	// RetryDelay is the delay between two consume try.
	// This take actions when the channel is disconnected for example.
	RetryDelay time.Duration
	// RequeueOnNackFn is a function that is called to have the requeue flag on a
	// nack response when the message consume is in error.
	// The default value is true is no function is set.
	RequeueOnNackFn func(d *amqp091.Delivery, err error) bool
}

// Client represents the AMQP client.
//
//go:generate mockgen -destination=./mocks/mock_Client.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/messagebus/amqp Client
type Client interface {
	// Connect will connect and create channels.
	Connect() error
	// Close will close all channels and connections.
	Close() error
	// Reconnect will handle the close and connect sequence.
	Reconnect() error
	// CancelAllConsumers will cancel all consumers.
	// This must be used in stop management.
	CancelAllConsumers() error
	// Publish will allow to publish a message.
	Publish(
		ctx context.Context,
		messageCfg *amqp091.Publishing,
		publishCfg *PublishConfigInput,
	) error
	// Consume will allow to consumer messages.
	// GetConsumeConfig is a function to allow the support of hot reloading the configuration.
	// Cb is a function that is called each time a message is handled.
	Consume(
		ctx context.Context,
		getConsumeCfg func() *ConsumeConfigInput,
		cb func(ctx context.Context, delivery *amqp091.Delivery) error,
	) error
	// Ping will check connections statuses.
	Ping() error
}

func New(
	logger log.Logger,
	cfgManager config.Manager,
	tracingSvc tracing.Service,
	signalHandlerSvc signalhandler.Client,
	metricsSvc metrics.Client,
) Client {
	return &amqpService{
		logger:           logger,
		cfgManager:       cfgManager,
		tracingSvc:       tracingSvc,
		signalHandlerSvc: signalHandlerSvc,
		metricsSvc:       metricsSvc,
		consumerTags:     []string{},
	}
}
