package app_service

import (
	"context"
	"errors"
	"mime/multipart"

	"first_move/config"
	"first_move/internal/app/domain_service"
	"first_move/internal/app/mutation"
)

func ParseTextChatFromForm(ctx context.Context, app *config.App, form *multipart.Form) error {
	app.Logger().Debug("Start loading csv...")

	// load csv
	fileHeaders := form.File["file"]
	if len(fileHeaders) != 1 {
		return errors.New("should contain 1 file")
	}

	fileHeader := fileHeaders[0]
	f, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, fileHeader.Size)
	_, err = f.Read(buf)

	// parse text chat (let openAI handle it)
	userBehaviors, err := domain_service.ParseTextChat(ctx, app, string(buf))
	if err != nil {
		return err
	}

	// insert rows to db
	tx, err := app.DB().Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = mutation.InsertUserBehavior(ctx, tx, userBehaviors)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
