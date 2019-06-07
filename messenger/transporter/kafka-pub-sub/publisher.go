package kafkapubsub

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	endpoint "github.com/go-kit/kit/endpoint"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// Publisher publishes messages to Google cloud pubsub
type Publisher struct {
	client *pubsub.Client
	enc    EncodeMessageFunc
	topic  *pubsub.Topic
}

// NewPublisher creates a publisher that will publish to the given topic
func NewPublisher(client *pubsub.Client, enc EncodeMessageFunc, topic string) *Publisher {
	return &Publisher{
		client: client,
		enc:    enc,
		topic:  client.Topic(topic),
	}
}

// Endpoint returns a useable endpoint for publishing messages on the publisher transport
func (p *Publisher) Endpoint() endpoint.Endpoint {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})

	return func(ctx context.Context, msg interface{}) (response interface{}, err error) {
		payload, err := p.enc(ctx, msg)
		if err != nil {
			return nil, err
		}

		_, err = p.topic.Publish(ctx, &pubsub.Message{Data: payload, PublishTime: time.Now()}).Get(ctx)

		return nil, err
	}
}

// Stop the publication
func (p *Publisher) Stop() error {
	p.topic.Stop()
	return nil
}
