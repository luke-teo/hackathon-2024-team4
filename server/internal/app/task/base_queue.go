package task

import (
	"go_chi_template/config"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type Queue struct {
	app *config.App
}

func NewQueue(app *config.App) *Queue {
	task := Queue{
		app: app,
	}
	return &task
}

func queueTask(app *config.App, task *asynq.Task) error {
	taskStatus, err := app.Queue().Client.Enqueue(
		task,
	)

	if err != nil {
		app.Logger().Error("Failed to queue task", zap.String("errorMessage", err.Error()))
		return err
	}

	app.Logger().Info(
		"Queued "+task.Type(),
		zap.String("taskType", task.Type()),
		zap.String("taskId", taskStatus.ID),
		zap.String("queue", taskStatus.Queue),
	)

	return nil
}
