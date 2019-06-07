package kafkapubsub

import (
	"context"
	"time"

	endpoint "github.com/go-kit/kit/endpoint"
	kafka "github.com/segmentio/kafka-go"
)

// Publisher publishes messages to Google cloud pubsub
type Publisher struct {
	config *Config
	enc    EncodeMessageFunc
}

type Config struct {
	Addr string
}

// NewPublisher creates a publisher that will publish to the given topic
func NewPublisher(config *Config, enc EncodeMessageFunc) *Publisher {
	return &Publisher{
		config: config,
		enc:    enc,
	}
}

// Endpoint returns a useable endpoint for publishing messages on the publisher transport
func (p *Publisher) Endpoint(topic string, partition int) endpoint.Endpoint {
	return func(ctx context.Context, msg interface{}) (response interface{}, err error) {
		conn, err := kafka.DialLeader(context.Background(), "tcp", p.config.Addr, topic, partition)
		if err != nil {
			return nil, err
		}
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		conn.WriteMessages(kafka.Message{Value: []byte(msg.(string))})
		conn.Close()
		return nil, err
	}
}

// Stop the publication
// func (p *Publisher) Stop() error {
// 	p.topic.Stop()
// 	return nil
// }
