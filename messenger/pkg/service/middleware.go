package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

type Middleware func(MessengerService) MessengerService

type loggingMiddleware struct {
	logger log.Logger
	next   MessengerService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next MessengerService) MessengerService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Subscriber(ctx context.Context, msg1 interface{}, attribute1 map[string]string) (msg interface{}, attribute map[string]string, err error) {
	defer func() {
		l.logger.Log("method", "Subscriber", "msg1", msg1, "", attribute1, "err", err)
	}()
	return l.next.Subscriber(ctx, msg1, attribute1)
}
