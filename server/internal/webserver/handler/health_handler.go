package handler

import (
	"net/http"
)

// this is a health check handler for the route "/health", it not use OpenAPI
func (h *Handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	// set status code to 200
	w.WriteHeader(http.StatusOK)
	// give a basic sucess message
	_, err := w.Write([]byte("Health Check Successful"))

	if err != nil {
		h.app.Logger().Error(err.Error())
	}
}
