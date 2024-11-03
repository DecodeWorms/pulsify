package pulsar

import (
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"
)

type PulsarClient struct {
	client pulsar.Client
}

func NewPulsarClient(url string) (*PulsarClient, error) {
	//Connecting to Pulsar Client
	fmt.Println("Connecting to Pulsar client")
	options := pulsar.ClientOptions{
		URL: url,
	}

	client, err := pulsar.NewClient(options)
	if err != nil {
		return nil, err
	}
	//Connected to Pulsar client successfully
	fmt.Println("Connected to Pulsar client successfully")

	return &PulsarClient{client: client}, nil
}

func (pc *PulsarClient) Close() {
	pc.client.Close()
}
