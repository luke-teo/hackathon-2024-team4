package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-jet/jet/v2/postgres"

	"github.com/getsentry/sentry-go"
)

func NewSentryProvider(env *EnvProvider) *sentryhttp.Handler {
	err := sentry.Init(sentry.ClientOptions{
		Environment:   env.sentryEnv,
		Dsn:           env.sentryDsn,
		EnableTracing: true,
		BeforeSendTransaction: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			// don't send for health check routes
			if event.Transaction == "" {
				return nil
			}

			if event.Transaction == "GET /health" {
				return nil
			}

			return event
		},
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	initializeDbTracing()

	// Create an instance of sentryhttp
	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	return sentryHandler
}

func initializeDbTracing() {
	postgres.SetQueryLogger(func(ctx context.Context, queryInfo postgres.QueryInfo) {
		now := time.Now()

		callerFile, callerLine, callerFunction := queryInfo.Caller()
		callerLog := fmt.Sprintf(
			"- Caller file: %s, line: %d, function: %s\n",
			callerFile,
			callerLine,
			callerFunction,
		)
		sqlStmt, _ := queryInfo.Statement.Sql()

		span := sentry.StartSpan(ctx, "db")
		span.StartTime = now.Add(-queryInfo.Duration)
		span.Description = sqlStmt

		// Depending on how the statement is executed, RowsProcessed is:
		//   - Number of rows returned for Query() and QueryContext() methods
		//   - RowsAffected() for Exec() and ExecContext() methods
		//   - Always 0 for Rows() method.
		span.SetData("Rows processed", strconv.Itoa(int(queryInfo.RowsProcessed)))
		span.SetData("Caller", callerLog)

		if queryInfo.Err != nil {
			span.SetData("Error", queryInfo.Err.Error())
		}

		// Do not log args to sentry, only for local debugging
		// sql, args := queryInfo.Statement.Sql()
		// fmt.Printf("- SQL: %s Args: %v \n", sql, args)
		span.Finish()

	})
}
