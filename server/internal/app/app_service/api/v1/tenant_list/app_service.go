package tenantlist

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
	_ oapi.GetApiV1TenantRequestObject,
) (oapi.GetApiV1TenantResponseObject, error) {
	tenants, err := multi.GetTenantsWithUserCount(ctx, app.DB())

	if err != nil {
		return nil, err
	}

	tenantDtos := []oapi.Tenant{}

	for _, tenant := range tenants {
		tenantDto := oapi.Tenant{
			Id:        strconv.FormatInt(tenant.ID, 10),
			Name:      tenant.Name,
			ShortCode: tenant.ShortCode,
			UserCount: &tenant.UserCount,
		}
		tenantDtos = append(tenantDtos, tenantDto)
	}

	return &oapi.GetApiV1Tenant200JSONResponse{
		Tenants: tenantDtos,
	}, nil
}
