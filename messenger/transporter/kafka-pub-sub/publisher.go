package kafkapubsub

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"

	endpoint1 "github.com/marcpar/gcp-pubsub-kafka/messenger/pkg/endpoint"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

//var logger log.Logger

// Publisher publishes messages to Google cloud pubsub
type Publisher struct {
	//config *Config
	topic string
	config *kafka.ConfigMap
	enc    EncodeMessageFunc
}

// NewPublisher creates a publisher that will publish to the given topic
func NewPublisher(config *kafka.ConfigMap, topic string, enc EncodeMessageFunc) *Publisher {
	return &Publisher{
		config: config,
		topic: topic,
		enc:    enc,
	}
}

// Endpoint returns a useable endpoint for publishing messages on the publisher transport
func (p *Publisher) Endpoint() endpoint.Endpoint {
	return func(ctx context.Context, msg interface{}) (response interface{}, err error) {

		producer, err := kafka.NewProducer(p.config)
		if err != nil {
			panic(err)
		}

		deliveryChan := make(chan kafka.Event)

		message := msg.(endpoint1.SubscriberResponse)
		attributes := message.Attributes[p.topic]
		value := []byte(message.Msg.([]uint8))

		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &attributes, Partition: kafka.PartitionAny},
			Value:          value,
		}, deliveryChan)

		e := <-deliveryChan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			panic(m.TopicPartition.Error)
		}
		close(deliveryChan)
		return nil, nil
	}
}
