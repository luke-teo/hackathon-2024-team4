package handler

import (
	"context"

	"first_move/generated/oapi"
)

func (h *Handler) GetScoresByUserID(
	ctx context.Context,
	request oapi.GetScoresByUserIDRequestObject,
) (oapi.GetScoresByUserIDResponseObject, error) {
	return oapi.GetScoresByUserID200JSONResponse{
		Scores: []oapi.Score{},
		UserId: "1",
	}, nil
}
