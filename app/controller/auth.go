package controller

import (
	"encoding/json"
	"errors"
	"go-go/app/helper"
	"go-go/app/service"
	"net/http"
)

type authController struct {
	userModel userModelInterface
}

func NewAuthController(userModel userModelInterface) *authController {
	return &authController{userModel: userModel}
}

// check is all fields are required
func (a *authController) RequiredForm(data ...string) error {
	for _, d := range data {
		if d == "" {
			return errors.New("all fields are required")
		}
	}
	return nil
}

// check authorization
func (a *authController) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data map[string]interface{}
	err := decoder.Decode(&data)
	if err != nil {
		helper.JSON(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{"error": err.Error()},
		)
		return
	}

	// check if all fields are required
	err = a.RequiredForm(data["email"].(string), data["password"].(string))
	if err != nil {
		helper.JSON(
			w,
			http.StatusBadRequest,
			map[string]interface{}{"error": err.Error()},
		)
		return
	}

	// check if user exist
	user, err := a.userModel.FindByEmail(data["email"].(string))
	if err != nil {
		helper.JSON(w, http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}
	if user.Password != data["password"] {
		helper.JSON(w, http.StatusUnauthorized, map[string]interface{}{"error": "invalid password"})
		return
	}
	token, err := service.EncrytData(
		map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	)

	// check if error
	if err != nil {
		helper.JSON(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{"error": err.Error()},
		)
		return
	}

	// return token and user data
	helper.JSON(
		w,
		http.StatusOK,
		map[string]interface{}{
			"token": token,
			"user": map[string]interface{}{
				"id":       user.ID,
				"username": user.Username,
				"email":    user.Email,
			},
		},
	)
}

func (a *authController) Register(w http.ResponseWriter, r *http.Request) {
	//
}
