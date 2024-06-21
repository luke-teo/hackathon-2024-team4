package worker

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"

	"first_move/config"
	"first_move/internal/app/task"
	"first_move/internal/worker/middleware"
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

	q := task.NewQueue(w.app)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.Use(middleware.LoggingMiddleware)

	// register handlers
	mux.HandleFunc(task.PARSE_TEXT_CHAT, q.ParseTextChat)

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
