package handler

import (
	"context"

	"first_move/generated/oapi"
	"first_move/internal/app/app_service"
)

func (h *Handler) PostUploadCsv(
	ctx context.Context,
	request oapi.PostUploadCsvRequestObject,
) (oapi.PostUploadCsvResponseObject, error) {
	f, err := request.Body.ReadForm(1 * 1024 * 1024)
	if err != nil {
		return nil, err
	}

	err = app_service.ParseTextChatFromForm(ctx, h.app, f)

	return nil, nil
}
