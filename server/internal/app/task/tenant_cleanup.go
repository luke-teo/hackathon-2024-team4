package task

import (
	"context"
	"encoding/json"
	"fmt"
	"go_chi_template/config"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

const TENANT_CLEANUP = "tenant:cleanup"

type TenantCleanupPayload struct {
	Reason string
}

func DispatchTenantCleanup(
	app *config.App,
	reason string,
) error {
	payload, err := json.Marshal(TenantCleanupPayload{
		Reason: reason,
	})

	if err != nil {
		return err
	}

	task := asynq.NewTask(TENANT_CLEANUP, payload)

	err = queueTask(app, task)

	if err != nil {
		return err
	}

	return nil
}

func (t *Queue) HandleTenantCleanup(ctx context.Context, task *asynq.Task) error {
	var payload TenantCleanupPayload

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf(
			"Failed to unmarshal %s: %v: %w",
			task.Type(),
			err,
			asynq.SkipRetry,
		)
	}

	// Usually you will execute app service here
	t.app.Logger().Info("Task Log", zap.String("Reason", payload.Reason))

	return nil
}
