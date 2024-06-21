package handler

import (
	"context"

	"first_move/generated/oapi"
	"first_move/internal/app/app_service"
)

func (h *Handler) GetScoresByUserID(
	ctx context.Context,
	request oapi.GetScoresByUserIDRequestObject,
) (oapi.GetScoresByUserIDResponseObject, error) {
	res, err := app_service.GetScoresByUserId(
		ctx,
		h.app,
		request.UserId,
		request.Params.StartDate.Time,
		request.Params.EndDate.Time,
	)
	if err != nil {
		return nil, err
	}

	return oapi.GetScoresByUserID200JSONResponse{
		UserId: request.UserId,
		Scores: res,
	}, nil
}
