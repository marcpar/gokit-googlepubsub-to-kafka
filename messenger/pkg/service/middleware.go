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

func (l loggingMiddleware) Subscriber(ctx context.Context, msg interface{}, attribute map[string]string) (err error) {
	defer func() {
		l.logger.Log("method", "Subscriber", "Message", msg, "err", err)
	}()
	return l.next.Subscriber(ctx, msg, attribute)
}
