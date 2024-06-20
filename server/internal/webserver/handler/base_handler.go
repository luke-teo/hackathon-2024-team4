package handler

import (
	"go_chi_template/config"
	"go_chi_template/generated/oapi"
)

type Handler struct {
	oapi.StrictServerInterface
	app *config.App
}

func NewHandler(app *config.App) *Handler {
	handler := Handler{
		app: app,
	}
	return &handler
}
