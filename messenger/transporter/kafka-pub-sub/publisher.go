package kafkapubsub

import (
	"context"
	"fmt"

	endpoint "github.com/go-kit/kit/endpoint"

	endpoint1 "github.com/marcpar/gcp-pubsub-kafka/messenger/pkg/endpoint"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// Publisher publishes messages to Google cloud pubsub
type Publisher struct {
	//config *Config

	config *kafka.ConfigMap
	enc    EncodeMessageFunc
}

type Config struct {
	Addr string
}

// NewPublisher creates a publisher that will publish to the given topic
func NewPublisher(config *kafka.ConfigMap, enc EncodeMessageFunc) *Publisher {
	return &Publisher{
		config: config,
		enc:    enc,
	}
}

func (p *Publisher) Endpoint() endpoint.Endpoint {
	return func(ctx context.Context, msg interface{}) (response interface{}, err error) {
		producer, err := kafka.NewProducer(p.config)
		if err != nil {
			panic(err)
		}
		defer producer.Close()
		go func() {
			for e := range producer.Events() {
				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
					} else {
						fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
					}
				}
			}
		}()
		message := msg.(endpoint1.SubscriberResponse)
		attributes := message.Attributes["topicName"]
		value := message.Msg.([]uint8)
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &attributes, Partition: kafka.PartitionAny},
			Value:          []byte(value),
		}, nil)

		producer.Flush(15 * 1000)
		return nil, err
	}

}

// Endpoint returns a useable endpoint for publishing messages on the publisher transport
// func (p *Publisher) Endpoint() endpoint.Endpoint {
// 	return func(ctx context.Context, msg interface{}) (response interface{}, err error) {
// 		sr := msg.(endpoint1.SubscriberResponse)
// 		attributes := sr.Attributes["topicName"]
// 		value := sr.Msg.([]uint8)
// 		conn, err := kafka.DialLeader(context.Background(), "tcp", p.config.Addr, attributes, 0)
// 		if err != nil {
// 			return nil, err
// 		}
// 		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
// 		conn.WriteMessages(kafka.Message{Value: value})
// 		conn.Close()
// 		return nil, err
// 	}
// }

// Stop the publication
// func (p *Publisher) Stop() error {
// 	p.topic.Stop()
// 	return nil
// }
