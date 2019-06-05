package service

import "context"

// GcpPubsubService describes the service.
type GcpPubsubService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Subscriber(ctx context.Context, s string) (rs string, err error)
}
