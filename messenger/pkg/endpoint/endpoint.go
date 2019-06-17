package endpoint

import (
	"context"

	"cloud.google.com/go/pubsub"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/marcpar/gcp-pubsub-kafka/messenger/pkg/service"
)

// SubscriberRequest collects the request parameters for the Subscriber method.
type SubscriberRequest struct {
	Msg        interface{}
	Attributes map[string]string
}

// SubscriberResponse collects the response parameters for the Subscriber method.
type SubscriberResponse struct {
	Msg        interface{}
	Attributes map[string]string
	Err        error
}

// MakeSubscriberEndpoint returns an endpoint that invokes Subscriber on the service.
func MakeSubscriberEndpoint(s service.MessengerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SubscriberRequest)
		msg, attr, err := s.Subscriber(ctx, req.Msg, req.Attributes)
		return SubscriberResponse{Msg: msg, Attributes: attr, Err: err}, nil
	}
}

// Failed implements Failer.
func (r SubscriberResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Subscriber implements Service. Primarily useful in a client.
func (e Endpoints) Subscriber(ctx context.Context, msg *pubsub.Message) (err error) {
	request := SubscriberRequest{
		Msg: msg,
	}
	response, err := e.SubscriberEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SubscriberResponse).Err
}
