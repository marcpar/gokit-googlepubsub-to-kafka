package service

import (
	"context"

	"cloud.google.com/go/pubsub"
)

// MessengerService describes the service.
type MessengerService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)

	Subscriber(ctx context.Context, msg *pubsub.Message) (err error)
}

type basicMessengerService struct {
}

// NewBasicMessengerService returns a naive, stateless implementation of MessengerService.
func NewBasicMessengerService() MessengerService {

	return &basicMessengerService{}
}

// New returns a MessengerService with all of the expected middleware wired in.
func New(middleware []Middleware) MessengerService {
	var svc MessengerService = NewBasicMessengerService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicMessengerService) Subscriber(ctx context.Context, msg *pubsub.Message) (err error) {
	// TODO implement the business logic of Subscriber
	return err
}
