package app_service

import (
	"context"
	"time"

	"first_move/config"
	"first_move/generated/oapi"
	"first_move/internal/app/domain_service"
	"first_move/internal/app/repository"
)

func GetScoresByUserId(
	ctx context.Context,
	app *config.App,
	userId string,
	startDate time.Time,
	endDate time.Time,
) ([]oapi.Score, error) {
	// query from db
	userBehaviors, err := repository.GetUserBehaviorByUserId(ctx, app.DB(), userId)
	if err != nil {
		return nil, err
	}

	// process user behaviors
	res, err := domain_service.ParseUserBehavior(userBehaviors, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return res, nil
}
