package handler

import (
	"context"
	"go_chi_template/generated/oapi"
	usercreate "go_chi_template/internal/app/app_service/api/v1/user_create"
	userdetail "go_chi_template/internal/app/app_service/api/v1/user_detail"
	userlist "go_chi_template/internal/app/app_service/api/v1/user_list"
)

func (h *Handler) PostApiV1User(
	ctx context.Context,
	request oapi.PostApiV1UserRequestObject,
) (oapi.PostApiV1UserResponseObject, error) {
	resp, err := usercreate.Handle(ctx, h.app, request)

	return resp, err
}

func (h *Handler) GetApiV1User(
	ctx context.Context,
	request oapi.GetApiV1UserRequestObject,
) (oapi.GetApiV1UserResponseObject, error) {
	resp, err := userlist.Handle(ctx, h.app, request)
	return resp, err
}

func (h *Handler) GetApiV1UserUserId(
	ctx context.Context,
	request oapi.GetApiV1UserUserIdRequestObject,
) (oapi.GetApiV1UserUserIdResponseObject, error) {
	resp, err := userdetail.Handle(ctx, h.app, request)
	return resp, err
}
