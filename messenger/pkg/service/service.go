package service

import (
	"context"
)

// MessengerService describes the service.
type MessengerService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)

	Subscriber(ctx context.Context, msg1 interface{}, attribute1 map[string]string) (msg interface{}, attribute map[string]string, err error)
}

type basicMessengerService struct{}

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

func (b *basicMessengerService) Subscriber(ctx context.Context, msg1 interface{}, attribute1 map[string]string) (msg interface{}, attribute map[string]string, err error) {
	// TODO implement the business logic of Subscriber
	// fmt.Println("hello")
	return msg1, attribute1, nil
}
