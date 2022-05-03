package amqpbusmessage

import (
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)

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
