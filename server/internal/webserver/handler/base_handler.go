package handler

import (
	"first_move/config"
	"first_move/generated/oapi"
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
