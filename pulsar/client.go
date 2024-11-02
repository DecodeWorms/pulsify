package pulsar

import (
	"pulsify/config"

	"github.com/apache/pulsar-client-go/pulsar"
)

type PulsarClient struct {
	client pulsar.Client
}

func NewPulsarClient(config *config.Config) (*PulsarClient, error) {
	options := pulsar.ClientOptions{
		URL: config.Url,
	}

	if config.Authentication {
		options.Authentication = pulsar.NewAuthenticationToken(config.Token)
	}

	client, err := pulsar.NewClient(options)
	if err != nil {
		return nil, err
	}

	return &PulsarClient{client: client}, nil
}

func (pc *PulsarClient) Close() {
	pc.client.Close()
}
