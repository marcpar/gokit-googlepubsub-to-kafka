package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	endpoint "github.com/marcpar/gcp-pubsub-kafka/messenger/pkg/endpoint"
	googlepubsub "github.com/marcpar/gcp-pubsub-kafka/messenger/transporter/google-pubsub"
)

func makeSubscriberHandler(client *pubsub.Client, topic string, subscription string, endpoints endpoint.Endpoints, options []googlepubsub.SubscriberOption) *googlepubsub.Subscriber {
	return googlepubsub.NewSubscriber(client, topic, subscription, endpoints.SubscriberEndpoint, decodeSubscriberSub, options...)
	//m.Handle("/subscriber", http1.NewServer(endpoints.SubscriberEndpoint, decodeSubscriberRequest, encodeSubscriberResponse, options...))
}

func decodeSubscriberSub(_ context.Context, msg interface{}) (interface{}, error) {

	m := msg.(*pubsub.Message)
	return endpoint.SubscriberRequest{Msg: m.Data, Attributes: m.Attributes}, nil
}

// func encodeSubscriberResponse(ctx context.Context, w, response interface{}) (err error) {
// 	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
// 		ErrorEncoder(ctx, f.Failed(), w)
// 		return nil
// 	}
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	err = json.NewEncoder(w).Encode(response)
// 	return
// }
