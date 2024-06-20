package userdetail

import (
	"context"
	"go_chi_template/config"
	"go_chi_template/generated/oapi"
	"go_chi_template/internal/app/repository/multi"
	"strconv"
)

func Handle(
	ctx context.Context,
	app *config.App,
	req oapi.GetApiV1UserUserIdRequestObject,
) (oapi.GetApiV1UserUserIdResponseObject, error) {
	userId, err := strconv.ParseInt(req.UserId, 10, 64)

	if err != nil {
		return nil, err
	}

	row, err := multi.GetUserWithTenantById(ctx, app.DB(), int(userId))

	if err != nil {
		return nil, err
	}

	userDto := oapi.User{
		Id:    strconv.FormatInt(row.User.ID, 10),
		Name:  row.User.Name,
		Email: &row.User.Email,
		Tenant: &oapi.Tenant{
			Id:        strconv.FormatInt(row.Tenant.ID, 10),
			Name:      row.Tenant.Name,
			ShortCode: row.Tenant.ShortCode,
		},
	}

	return &oapi.GetApiV1UserUserId200JSONResponse{User: userDto}, nil
}
