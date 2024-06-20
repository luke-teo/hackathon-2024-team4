package handler

import (
	"context"
	"go_chi_template/generated/oapi"
	tenantcreate "go_chi_template/internal/app/app_service/api/v1/tenant_create"
	tenantdetail "go_chi_template/internal/app/app_service/api/v1/tenant_detail"
	tenantlist "go_chi_template/internal/app/app_service/api/v1/tenant_list"
	tenantupdate "go_chi_template/internal/app/app_service/api/v1/tenant_update"
)

func (h *Handler) GetApiV1Tenant(
	ctx context.Context,
	request oapi.GetApiV1TenantRequestObject,
) (oapi.GetApiV1TenantResponseObject, error) {
	resp, err := tenantlist.Handle(ctx, h.app, request)
	return resp, err
}

func (h *Handler) PostApiV1Tenant(
	ctx context.Context,
	request oapi.PostApiV1TenantRequestObject,
) (oapi.PostApiV1TenantResponseObject, error) {
	resp, err := tenantcreate.Handle(ctx, h.app, request)

	return resp, err
}

func (h *Handler) GetApiV1TenantTenantId(
	ctx context.Context,
	request oapi.GetApiV1TenantTenantIdRequestObject,
) (oapi.GetApiV1TenantTenantIdResponseObject, error) {
	resp, err := tenantdetail.Handle(ctx, h.app, request)
	return resp, err
}

func (h *Handler) PatchApiV1TenantTenantId(
	ctx context.Context,
	request oapi.PatchApiV1TenantTenantIdRequestObject,
) (oapi.PatchApiV1TenantTenantIdResponseObject, error) {
	resp, err := tenantupdate.Handle(ctx, h.app, request)
	return resp, err
}
