package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type internalServerErrorResponse struct {
	ErrorCode string `json:"errorCode"`
}

func ErrorResponseHandler(
	l *zap.Logger,
	w http.ResponseWriter,
	r *http.Request,
	err error,
	errorCode string,
) {
	// Log error
	l.Error("Internal Server Error: ",
		zap.String("proto", r.Proto),
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("time", time.Now().String()),
		zap.Error(err),
	)
	// Return general response
	w.WriteHeader(http.StatusInternalServerError)
	res := &internalServerErrorResponse{
		ErrorCode: errorCode,
	}

	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Panic(err)
	}
}
