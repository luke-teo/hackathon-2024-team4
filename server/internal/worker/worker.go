package worker

import (
	"context"
	"go_chi_template/config"
	"go_chi_template/internal/app/task"
	"go_chi_template/internal/worker/middleware"
	"log"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type Worker struct {
	app *config.App
}

func NewWorker(app *config.App) *Worker {
	return &Worker{
		app: app,
	}
}

const (
	concurrentWorkers  = 3
	criticalQueueLevel = 6
	defaultQueueLevel  = 3
	lowQueueLevel      = 1
)

func (w *Worker) Start() {
	redisOpts := w.app.Redis().Options()

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:      redisOpts.Addr,
			Password:  redisOpts.Password,
			TLSConfig: redisOpts.TLSConfig,
		},
		asynq.Config{
			Concurrency:  concurrentWorkers,
			ErrorHandler: asynq.ErrorHandlerFunc(w.handleError),
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.Use(middleware.LoggingMiddleware)
	t := task.NewQueue(w.app)

	// register handlers
	mux.HandleFunc(task.TENANT_CLEANUP, t.HandleTenantCleanup)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("Failed to start worker: %v", err)
	}
}

func (w *Worker) handleError(ctx context.Context, task *asynq.Task, err error) {
	taskID, _ := asynq.GetTaskID(ctx)
	taskRetryCount, _ := asynq.GetRetryCount(ctx)
	taskType := task.Type()
	taskPayload := task.Payload()

	w.app.Logger().Error(
		"Error handling task",
		zap.String("taskType", taskType),
		zap.String("taskID", taskID),
		zap.Int("taskRetryCount", taskRetryCount),
		zap.String("taskPayload", string(taskPayload)),
		zap.Error(err),
	)
}
