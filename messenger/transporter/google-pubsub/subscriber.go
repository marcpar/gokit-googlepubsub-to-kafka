package googlepubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	endpoint "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
)

type Handler interface {
	Serve()
	rcv(context.Context, *pubsub.Message)
	Stop()
}

// Subscriber receives messages from Google cloud pubsub
type Subscriber struct {
	logger           log.Logger
	client           *pubsub.Client
	subscription     *pubsub.Subscription
	endpoint         endpoint.Endpoint
	dec              DecodeMessageFunc
	responseEndpoint endpoint.Endpoint
	errorEndpoint    endpoint.Endpoint
	topic            *pubsub.Topic
}

// SubscriberOption sets an optional parameter for clients
type SubscriberOption func(*Subscriber)

// WithLogger specifies the logger to use
func WithLogger(logger log.Logger) SubscriberOption {
	return func(s *Subscriber) {
		s.logger = log.With(logger, "topic", s.topic, "subscription", s.subscription.ID())
	}
}

// WithResponseEndpoint specifies an endpoint to use for sending response returned by the subscriber endpoint
func WithResponseEndpoint(endpoint endpoint.Endpoint) SubscriberOption {
	return func(s *Subscriber) {
		s.responseEndpoint = endpoint
	}
}

// WithErrorEndpoint specifies an endpoint to use for sending errors returned by the subscriber endpoint
func WithErrorEndpoint(endpoint endpoint.Endpoint) SubscriberOption {
	return func(s *Subscriber) {
		s.errorEndpoint = endpoint
	}
}

// NewSubscriber create a subscription of the endpoint to the given event
func NewSubscriber(client *pubsub.Client, topicName string, subscription string, e endpoint.Endpoint, dec DecodeMessageFunc, options ...SubscriberOption) *Subscriber {
	//subscriptionID := topicName
	s := &Subscriber{
		client:       client,
		subscription: client.Subscription(subscription),
		endpoint:     e,
		dec:          dec,
		topic:        client.Topic(topicName),
		logger:       log.NewNopLogger(),
	}

	for _, option := range options {
		option(s)
	}

	return s
}

// Serve begins listening for messages and passing them to the endpoint:
func (s *Subscriber) Serve() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	exists, err := s.subscription.Exists(ctx)
	if err != nil {
		return err
	}

	if !exists {
		if s.subscription, err = s.client.CreateSubscription(ctx, s.subscription.ID(), pubsub.SubscriptionConfig{Topic: s.topic}); err != nil {
			return err
		}
	}

	err = s.subscription.Receive(ctx, s.rcv)

	if err == nil {
		s.logger.Log("status", "listening")
	}

	return err
}

func (s *Subscriber) rcv(ctx context.Context, msg *pubsub.Message) {

	defer func() {

		if r := recover(); r != nil {
			s.logger.Log("error", r)
			msg.Nack()
		} else {
			msg.Ack()
		}
	}()

	payload, err := s.dec(ctx, msg)

	if err != nil {
		s.logger.Log("error", err)
		return
	}

	response, err := s.endpoint(ctx, payload)

	if err != nil {
		if s.errorEndpoint == nil {
			s.logger.Log("error", err)
			return
		}

		if _, err := s.errorEndpoint(ctx, err); err != nil {
			s.logger.Log("error", err)
			return
		}
	}

	if response != nil {
		if s.responseEndpoint == nil {
			s.logger.Log("error", "Response not delivered. Response endpoint not configured.")
		}

		if _, err := s.responseEndpoint(ctx, response); err != nil {
			s.logger.Log("error", err)
			return
		}
	}
}

// Stop can gracefully stops the subscriber
func (s *Subscriber) Stop() error {

	err := s.subscription.Delete(context.Background())
	if err == nil {
		s.logger.Log("status", "stopped")
	}
	return err
}
