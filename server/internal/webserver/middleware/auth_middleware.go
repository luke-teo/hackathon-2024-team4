package middleware

import (
	"errors"
	"fmt"
	"go_chi_template/config"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func validateHeader(_app *config.App, bearerToken string) error {
	// will parse and validate signature (including issued at and expiration)
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		// only HMAC is alllowed
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// this needs to be set
		return []byte{}, nil
	})

	if err != nil {
		return err
	}

	sub, err := token.Claims.GetSubject()

	if err != nil {
		return err
	}

	// TODO: set your expected JWT subject (e.g. app client ID, user ID)
	expectedSubject := ""

	if sub != expectedSubject {
		return errors.New("Invalid JWT Subject")
	}

	return nil
}

func NewAuthMiddleware(app *config.App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			bearerToken := r.Header.Get("authorization")

			err := validateHeader(app, bearerToken)

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
