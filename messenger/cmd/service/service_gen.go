// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	endpoint "github.com/marcpar/gcp-pubsub-kafka/messenger/pkg/endpoint"
	http1 "github.com/marcpar/gcp-pubsub-kafka/messenger/pkg/http"
	service "github.com/marcpar/gcp-pubsub-kafka/messenger/pkg/service"
	googlepubsub "github.com/marcpar/gcp-pubsub-kafka/messenger/transporter/google-pubsub"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	initGooglePubSubHandler(endpoints.SubscriberEndpoint, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"Subscriber": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Subscriber", logger))}}
	return options
}

func defaultGooglePubSubOptions(logger log.Logger, tracer opentracinggo.Tracer, endpoint endpoint1.Endpoint) map[string][]googlepubsub.SubscriberOption {
	options := map[string][]googlepubsub.SubscriberOption{
		"Subscriber": {googlepubsub.WithErrorEndpoint(endpoint)}}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["Subscriber"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Subscriber")), endpoint.InstrumentingMiddleware(duration.With("method", "Subscriber"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Subscriber"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}