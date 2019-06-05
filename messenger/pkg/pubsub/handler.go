package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	endpoint "github.com/marcpar/gcp-pubsub-kafka/messenger/pkg/endpoint"
	googlepubsub1 "github.com/marcpar/gcp-pubsub-kafka/messenger/transporter/google-pubsub"
)

func makeSubscriberHandler(client *pubsub.Client, topic string, subscription string, endpoints endpoint.Endpoints, options []googlepubsub1.SubscriberOption) {
	googlepubsub1.NewSubscriber(client, topic, subscription, endpoints.SubscriberEndpoint, decodeSubscriberSub, options...)
	//m.Handle("/subscriber", http1.NewServer(endpoints.SubscriberEndpoint, decodeSubscriberRequest, encodeSubscriberResponse, options...))
}

func decodeSubscriberSub(_ context.Context, b []byte) (interface{}, error) {
	message := endpoint.SubscriberRequest{}
	fmt.Println(message)
	//err := json.NewDecoder(b).Decode(&b)
	return message, nil
}
