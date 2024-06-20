package handler

import (
	"context"
	"go_chi_template/generated/oapi"
	departmentlist "go_chi_template/internal/app/app_service/api/v1/department_list"
)

func (h *Handler) GetApiV1Department(
	ctx context.Context,
	request oapi.GetApiV1DepartmentRequestObject,
) (oapi.GetApiV1DepartmentResponseObject, error) {
	resp, err := departmentlist.Handle(ctx, h.app, request)
	return resp, err
}
