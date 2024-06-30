package controller

import (
	"go-go/app/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	helper.JSON(
		w,
		http.StatusOK,
		map[string]interface{}{
			"message": "Welcome to Go-Go",
		},
	)
}
