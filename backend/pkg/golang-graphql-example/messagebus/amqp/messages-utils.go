package amqpbusmessage

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
)

// CreateJSONPublishingMessage will create a publish message from data with a json encoded content and content
// type data set to "application/json".
func CreateJSONPublishingMessage(data interface{}) (*amqp091.Publishing, error) {
	// Encode it
	b, err := json.Marshal(data)
	// Check error
	if err != nil {
		return nil, err
	}

	// Create message
	return &amqp091.Publishing{
		ContentType: "application/json",
		Body:        b,
	}, nil
}

// ParseJSONMessage will parse a message as json object if possible.
func ParseJSONMessage(res interface{}, input *amqp091.Delivery) error {
	// Check content type
	if input.ContentType != "application/json" {
		return ErrMessageNotJSON
	}

	// Try to unmarshal
	err := json.Unmarshal(input.Body, res)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Default
	return nil
}
