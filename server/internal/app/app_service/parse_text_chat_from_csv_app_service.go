package app_service

import (
	"context"
	"os"

	"first_move/config"
	"first_move/internal/app/domain_service"
)

func ParseTextChatFromCSV(ctx context.Context, app *config.App, path string) error {
	app.Logger().Debug("Start loading csv...")

	// load csv
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	app.Logger().Debug("Finished loading csv!")

	// parse text chat (let openAI handle it)
	_, err = domain_service.ParseTextChat(ctx, app, string(buf))
	if err != nil {
		return err
	}

	return nil
}
