package tenantcreate

import (
	"context"
	"fmt"
	"go_chi_template/config"
	"go_chi_template/generated/oapi"
	domainservice "go_chi_template/internal/app/domain_service"
	"go_chi_template/internal/app/mutation"
	validatetenantcreate "go_chi_template/internal/app/util_service/validate_tenant_create"

	"go.uber.org/zap"
)

func Handle(
	ctx context.Context,
	app *config.App,
	req oapi.PostApiV1TenantRequestObject,
) (oapi.PostApiV1TenantResponseObject, error) {
	errResp, err := validatetenantcreate.Handle(ctx, app, req.Body)

	if err != nil {
		return nil, err
	}

	if errResp != nil {
		return errResp, nil
	}

	newTenant := domainservice.NewTenant(req.Body.Name, req.Body.ShortCode)
	insertedTenant, err := mutation.InsertTenant(ctx, app.DB(), newTenant)

	if err != nil {
		app.Logger().With(
			zap.String("message", "insert tenant failed"),
			zap.String("name", req.Body.Name),
			zap.String("shortCode", req.Body.ShortCode),
		)

		return nil, err
	}

	resp := oapi.PostApiV1Tenant200JSONResponse{
		Tenant: oapi.Tenant{
			Id:        fmt.Sprint(insertedTenant.ID),
			Name:      insertedTenant.Name,
			ShortCode: insertedTenant.ShortCode,
		},
	}

	return &resp, nil
}
