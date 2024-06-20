package usercreate

import (
	"context"
	"go_chi_template/config"
	"go_chi_template/generated/oapi"
	domainservice "go_chi_template/internal/app/domain_service"
	"go_chi_template/internal/app/mutation"
	"go_chi_template/internal/app/repository/single"
	validateusercreate "go_chi_template/internal/app/util_service/validate_user_create"
	"strconv"
)

func Handle(
	ctx context.Context,
	app *config.App,
	req oapi.PostApiV1UserRequestObject,
) (oapi.PostApiV1UserResponseObject, error) {
	respDto, err := validateusercreate.Handle(ctx, app, req.Body)

	if err != nil {
		return nil, err
	}

	if respDto != nil {
		return respDto, nil
	}

	tenantId, err := strconv.ParseInt(req.Body.TenantId, 10, 64)

	if err != nil {
		return nil, err
	}

	tenant, err := single.GetTenantById(ctx, app.DB(), tenantId)

	if err != nil {
		return nil, err
	}

	user := domainservice.NewUser(tenant, req.Body.Name, req.Body.Email)
	insertedUser, err := mutation.InsertUser(ctx, app.DB(), user)

	if err != nil {
		return nil, err
	}

	role := "admin"

	userDto := oapi.User{
		Id:    strconv.FormatInt(insertedUser.ID, 10),
		Name:  insertedUser.Name,
		Email: &insertedUser.Email,
		Role:  &role,
	}

	return &oapi.PostApiV1User200JSONResponse{User: userDto}, err
}
