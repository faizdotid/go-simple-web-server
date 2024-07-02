package controller

import (
	"encoding/json"
	"errors"
	"go-go/app/helper"
	"go-go/app/model"
	"go-go/app/service"
	"net/http"
)

// auth controller struct
type authController struct {
	userModel userModelInterface
}

// new auth controller
func NewAuthController(userModel userModelInterface) *authController {
	return &authController{
		userModel: userModel,
	}
}

// check is all fields are required
func (a *authController) RequiredForm(r *http.Request, required ...string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, errors.New("invalid request body")
	}
	if len(data) == 0 {
		return nil, errors.New("request body is empty")
	}
	for _, v := range data {
		if v == nil {
			return nil, errors.New("all fields are required")
		}
	}
	return data, nil
}

// login user
func (a *authController) Login(w http.ResponseWriter, r *http.Request) {
	data, err := a.RequiredForm(r, "email", "password")
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

// register user
func (a *authController) Register(w http.ResponseWriter, r *http.Request) {
	data, err := a.RequiredForm(r, "username", "email", "password")
	if err != nil {
		helper.JSON(
			w,
			http.StatusBadRequest,
			map[string]interface{}{"error": err.Error()},
		)
		return
	}

	// check if user exist
	_, err = a.userModel.FindByEmail(data["email"].(string))
	if err == nil {
		helper.JSON(w, http.StatusConflict, map[string]interface{}{"error": "user already exist"})
		return
	}

	// create user
	user, err := a.userModel.Create(model.User{
		Username: data["username"].(string),
		Email:    data["email"].(string),
		Password: data["password"].(string),
	})

	// check if error
	if err != nil {
		helper.JSON(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{"error": err.Error()},
		)
		return
	}
	token, err := service.EncrytData(
		map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	)
	if err != nil {
		helper.JSON(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{"error": err.Error()},
		)
		return
	}

	// return user data
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
