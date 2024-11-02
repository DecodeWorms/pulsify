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

func (c *Consumer) ReceiveMessage() {
	for {
		msg, err := c.consumer.Receive(context.Background())
		if err != nil {
			log.Printf("Failed to receive message: %v", err)
			continue
		}

		log.Printf("Received message: %s", string(msg.Payload()))
		c.consumer.Ack(msg)
	}
}

func (c *Consumer) Close() {
	c.consumer.Close()
}
