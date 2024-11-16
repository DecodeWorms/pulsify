package pulsar

import (
	"context"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

type Consumer struct {
	consumer pulsar.Consumer
}

func (pc *PulsarClient) CreateConsumer(topic, subscription string) (*Consumer, error) {
	consumer, err := pc.client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: subscription,
		Type:             pulsar.Shared,
	})
	if err != nil {
		return nil, err
	}

	return &Consumer{consumer: consumer}, nil
}

func (c *Consumer) ReceiveMessage() (interface{}, error) {
	// Retrieve the most recent message from Pulsar
	msg, err := c.consumer.Receive(context.Background())
	if err != nil {
		log.Printf("Failed to receive message: %v", err)
		return "", err
	}

	// Process the message
	payload := string(msg.Payload())
	log.Printf("Received message: %s", payload)

	// Acknowledge the message
	c.consumer.Ack(msg)

	// Return the payload of the last message received
	return payload, nil

}

func (c *Consumer) Close() {
	c.consumer.Close()
}
