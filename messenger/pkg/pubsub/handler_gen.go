package pubsub

import (
	"cloud.google.com/go/pubsub"
	endpoint "github.com/marcpar/gcp-pubsub-kafka/messenger/pkg/endpoint"
	googlepubsub "github.com/marcpar/gcp-pubsub-kafka/messenger/transporter/google-pubsub"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewGCPPubSubHandler(client *pubsub.Client, topic string, subscription string, endpoints endpoint.Endpoints, options map[string][]googlepubsub.SubscriberOption) {
	//m := http1.NewServeMux()
	makeSubscriberHandler(client, topic, subscription, endpoints, options["Subscriber"])

}
