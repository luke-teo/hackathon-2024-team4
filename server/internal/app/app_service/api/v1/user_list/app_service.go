package userlist

import (
	"context"
	"go_chi_template/config"
	"go_chi_template/generated/oapi"
	"go_chi_template/internal/app/repository/single"
	"strconv"
)

func Handle(
	ctx context.Context,
	app *config.App,
	req oapi.GetApiV1UserRequestObject,
) (oapi.GetApiV1UserResponseObject, error) {
	tenantId, err := strconv.ParseInt(req.Body.TenantId, 10, 64)

	if err != nil {
		return nil, err
	}

	users, err := single.GetUsersByTenantId(ctx, app.DB(), tenantId)

	if err != nil {
		return nil, err
	}

	userDtos := []oapi.User{}

	for _, user := range users {
		userDto := oapi.User{
			Id:    strconv.FormatInt(user.ID, 10),
			Name:  user.Name,
			Email: user.Name,
		}
		userDtos = append(userDtos, userDto)
	}

	return &oapi.GetApiV1User200JSONResponse{
		Users: userDtos,
	}, nil
}
