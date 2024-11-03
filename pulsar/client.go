package pulsar

import (
	"github.com/apache/pulsar-client-go/pulsar"
)

type PulsarClient struct {
	client pulsar.Client
}

func NewPulsarClient(url, token string, auth bool) (*PulsarClient, error) {
	options := pulsar.ClientOptions{
		URL: url,
	}

	if auth {
		options.Authentication = pulsar.NewAuthenticationToken(token)
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
