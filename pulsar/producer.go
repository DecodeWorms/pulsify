package pulsar

import (
	"context"
	"encoding/json"
	"log"

	"github.com/DecodeWorms/pulsify/model"
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

func (p *Producer) SendMessage(msg model.VerifyEmail) error {
	var ctx = context.Background()

	//Marshal the msg to json
	m, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = p.producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: m,
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
