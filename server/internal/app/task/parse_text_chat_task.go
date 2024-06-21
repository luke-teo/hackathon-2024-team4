package task

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"

	"first_move/config"
	"first_move/internal/app/app_service"
)

const PARSE_TEXT_CHAT = "task:parse_text_chat"

type ParseTextChatPayload struct {
	Filepath string
}

func DispatchParseTextChatTask(
	app *config.App,
	filepath string,
) error {
	payload, err := json.Marshal(ParseTextChatPayload{
		Filepath: filepath,
	})
	if err != nil {
		return err
	}

	task := asynq.NewTask(PARSE_TEXT_CHAT, payload)

	err = queueTask(app, task)
	if err != nil {
		return err
	}

	return nil
}

func (t *Queue) ParseTextChat(
	ctx context.Context,
	task *asynq.Task,
) error {
	var payload ParseTextChatPayload

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf(
			"Failed to unmarshal %s: %v: %w",
			task.Type(),
			err,
			asynq.SkipRetry,
		)
	}

	t.app.Logger().Info("Start executing ParseTextChat job...")
	fmt.Println(payload)

	return app_service.ParseTextChatFromCSV(ctx, t.app, payload.Filepath)
}
