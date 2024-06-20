package tenantupdate

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
	req oapi.PatchApiV1TenantTenantIdRequestObject,
) (oapi.PatchApiV1TenantTenantIdResponseObject, error) {
	tenantId, err := strconv.ParseInt(req.Body.TenantId, 10, 64)

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

	respDto := oapi.PatchApiV1TenantTenantId200JSONResponse{
		Tenant: &tenantDto,
	}

	return &respDto, nil
}
