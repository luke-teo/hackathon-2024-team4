package middleware

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/hibiken/asynq"
)

func LoggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		hub := sentry.GetHubFromContext(ctx)
		if hub == nil {
			hub = sentry.CurrentHub().Clone()
			ctx = sentry.SetHubOnContext(ctx, hub)
		}
		options := []sentry.SpanOption{
			sentry.WithOpName("queue.worker"),
			sentry.WithTransactionSource(sentry.SourceURL),
		}

		transaction := sentry.StartTransaction(ctx,
			t.Type(),
			options...,
		)
		defer transaction.Finish()

		err := h.ProcessTask(ctx, t)
		if err != nil {
			transaction.Status = sentry.SpanStatusInternalError
			sentry.CaptureException(err)
			return err
		}

		transaction.Status = sentry.SpanStatusOK

		return nil
	})
}
