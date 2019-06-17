package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	endpoint "github.com/marcpar/gcp-pubsub-kafka/messenger/pkg/endpoint"
	googlepubsub "github.com/marcpar/gcp-pubsub-kafka/messenger/transporter/google-pubsub"
	kafkapubsub "github.com/marcpar/gcp-pubsub-kafka/messenger/transporter/kafka-pub-sub"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func makeSubscriberHandler(client *pubsub.Client, topic string, subscription string, endpoints endpoint.Endpoints, options []googlepubsub.SubscriberOption) *googlepubsub.Subscriber {
	return googlepubsub.NewSubscriber(client, topic, subscription, endpoints.SubscriberEndpoint, decodeSubscriberSub, options...)
}

func decodeSubscriberSub(_ context.Context, msg interface{}) (interface{}, error) {
	m := msg.(*pubsub.Message)
	return endpoint.SubscriberRequest{Msg: m.Data, Attributes: m.Attributes}, nil
}

func makePublisherHandler(config *kafka.ConfigMap, topic string) *kafkapubsub.Publisher {
	return kafkapubsub.NewPublisher(config, topic, encodeSubscriberPub)
}

func encodeSubscriberPub(_ context.Context, response interface{}) ([]byte, error) {
	topic := response.(string)
	return []byte(topic), nil
}
