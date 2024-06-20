package tenantdetail

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
	req oapi.GetApiV1TenantTenantIdRequestObject,
) (oapi.GetApiV1TenantTenantIdResponseObject, error) {
	tenantId, err := strconv.ParseInt(req.TenantId, 10, 64)

	if err != nil {
		return nil, err
	}

	tenant, err := single.GetTenantById(ctx, app.DB(), tenantId)

	if err != nil {
		return nil, err
	}

	tenantDto := oapi.Tenant{
		Id:        strconv.FormatInt(tenant.ID, 10),
		Name:      tenant.Name,
		ShortCode: tenant.ShortCode,
	}

	respDto := oapi.GetApiV1TenantTenantId200JSONResponse{
		Tenant: tenantDto,
	}

	return &respDto, nil
}
