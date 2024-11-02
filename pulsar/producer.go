package pulsar

import (
	"context"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

type Producer struct {
	producer pulsar.Producer
}

func (pc *PulsarClient) CreateProducer(topic string) (*Producer, error) {
	producer, err := pc.client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		return nil, err
	}

	return &Producer{producer: producer}, nil
}

func (p *Producer) SendMessage(msg string) error {
	var ctx = context.Background()

	_, err := p.producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: []byte(msg),
	})
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}

	log.Printf("Message sent: %s", msg)
	return nil
}

func (p *Producer) Close() {
	p.producer.Close()
}
