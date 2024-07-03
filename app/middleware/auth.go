package middleware

import (
	"context"
	"go-go/app/helper"
	"go-go/app/service"
	"net/http"
	"strings"
)

type User string

func Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			helper.JSON(
				w,
				http.StatusUnauthorized,
				map[string]interface{}{"error": "Unauthorized"},
			)
			return
		}
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			next(w, r)
			return
		}
		token := strings.Split(auth, " ")[1]
		if token == "" {
			next(w, r)
			return
		}
		data, err := service.DecryptData(token)
		if err != nil {
			helper.JSON(
				w,
				http.StatusUnauthorized,
				map[string]interface{}{"error": err.Error()},
			)
			return
		}
		ctx := context.WithValue(r.Context(), User("user"), data)
		next(w, r.WithContext(ctx))
	}
}
