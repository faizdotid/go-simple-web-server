package controller

import (
	"go-go/app/helper"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	helper.JSON(
		w,
		http.StatusOK,
		map[string]interface{}{"message": "ping"},
	)
}
