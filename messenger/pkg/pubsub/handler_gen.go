package pubsub

import (
	"cloud.google.com/go/pubsub"
	endpoint "github.com/marcpar/gcp-pubsub-kafka/messenger/pkg/endpoint"
	googlepubsub "github.com/marcpar/gcp-pubsub-kafka/messenger/transporter/google-pubsub"
	kafkapubsub "github.com/marcpar/gcp-pubsub-kafka/messenger/transporter/kafka-pub-sub"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewGCPPubSubHandler(client *pubsub.Client, topic string, subscription string, endpoints endpoint.Endpoints, options map[string][]googlepubsub.SubscriberOption) *googlepubsub.Subscriber {
	return makeSubscriberHandler(client, topic, subscription, endpoints, options["Subscriber"])
}

// NewKafkaPubSubHandler lorem ipsum
func NewKafkaPubSubHandler(config *kafka.ConfigMap, topic string) *kafkapubsub.Publisher {
	return makePublisherHandler(config, topic)
}
