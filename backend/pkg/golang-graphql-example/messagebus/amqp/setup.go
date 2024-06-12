package amqpbusmessage

import (
	"emperror.dev/errors"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (as *amqpService) ExtraSetup(input *ExtraSetupInput) error {
	// Get a valid channel
	chann := as.consumerChannel
	// Check if it isn't opened or null
	if chann == nil || !chann.IsClosed() {
		chann = as.publisherChannel
	}
	// Check if it isn't opened or null
	if chann == nil || !chann.IsClosed() {
		return errors.WithStack(ErrNoActiveChannelFound)
	}

	// Build configuration
	cfg := &config.AMQPConfig{
		Exchanges:  input.Exchanges,
		Queues:     input.Queues,
		QueueBinds: input.QueueBinds,
	}

	return as.setup(cfg, chann)
}

func (*amqpService) setup(cfg *config.AMQPConfig, chann *amqp.Channel) error {
	// Declare exchanges
	// Loop over exchange configurations
	for _, it := range cfg.Exchanges {
		// Declare it
		err := chann.ExchangeDeclare(
			it.Name,
			it.Type,
			it.Durable,
			it.AutoDelete,
			it.Internal,
			it.NoWait,
			it.ExtraArgs,
		)
		// Check error
		if err != nil {
			return errors.Wrap(err, "error in exchange declaration")
		}
	}

	// Declare queues
	// Loop over queue configurations
	for _, it := range cfg.Queues {
		// Declare it
		_, err := chann.QueueDeclare(
			it.Name,
			it.Durable,
			it.AutoDelete,
			it.Exclusive,
			it.NoWait,
			it.ExtraArgs,
		)
		// Check error
		if err != nil {
			return errors.Wrap(err, "error in queue declaration")
		}
	}

	// Bind queues
	// Loop over bind queue configurations
	for _, it := range cfg.QueueBinds {
		// Bind it
		err := chann.QueueBind(
			it.Name,
			it.Key,
			it.Exchange,
			it.NoWait,
			it.ExtraArgs,
		)
		// Check error
		if err != nil {
			return errors.Wrap(err, "error in queue bind")
		}
	}

	// Default
	return nil
}
