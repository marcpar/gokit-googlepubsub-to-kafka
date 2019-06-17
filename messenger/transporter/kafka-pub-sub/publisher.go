package kafkapubsub

import (
	"context"
	"fmt"

	endpoint "github.com/go-kit/kit/endpoint"

	endpoint1 "github.com/marcpar/gcp-pubsub-kafka/messenger/pkg/endpoint"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

//var logger log.Logger

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

// Endpoint returns a useable endpoint for publishing messages on the publisher transport
func (p *Publisher) Endpoint() endpoint.Endpoint {
	return func(ctx context.Context, msg interface{}) (response interface{}, err error) {

		producer, err := kafka.NewProducer(p.config)
		if err != nil {
			panic(err)
		}

		deliveryChan := make(chan kafka.Event)

		message := msg.(endpoint1.SubscriberResponse)
		attributes := message.Attributes["topicName"]
		value := []byte(message.Msg.([]uint8))

		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &attributes, Partition: kafka.PartitionAny},
			Key:	[]byte("test"),
			Value:          value,
			Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
		}, deliveryChan)

		// if err != nil {
		// 	panic(err)
		// }

		e := <-deliveryChan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
		close(deliveryChan)

		fmt.Println("attributes", attributes)
		fmt.Println("string(value)", string(value))

		return nil, nil
	}
}
