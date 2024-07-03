package controller

import (
	"go-go/app/helper"
	"go-go/app/model"
	"go-go/app/service"
	"net/http"
	"strings"
)

type userModelInterface interface {
	FindAll() ([]model.User, error)
	FindByEmail(email string) (model.User, error)
	FindById(id int) (model.User, error)
	Create(data model.User) (model.User, error)
	Delete(id int) (model.User, error)
	Update(id int, data model.User) (model.User, error)
}

type userControllers struct {
	userModel userModelInterface
}

func NewUserController(userModel userModelInterface) *userControllers {
	return &userControllers{userModel: userModel}
}

func (u *userControllers) Index(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	bearer := strings.Split(auth, " ")[1]
	user, err := service.DecryptData(bearer)
	if err != nil {
		helper.JSON(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{"error": err.Error()},
		)
		return
	}
	helper.JSON(
		w,
		http.StatusOK,
		map[string]interface{}{"user": user},
	)
}

func (u *userControllers) Posts(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	bearer := strings.Split(auth, " ")[1]
	user, err := service.DecryptData(bearer)
	if err != nil {
		helper.JSON(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{"error": err.Error()},
		)
		return
	}
	posts, err := u.userModel.FindById(int(user["id"].(float64)))
	if err != nil {
		helper.JSON(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{"error": err.Error()},
		)
		return
	}
	helper.JSON(
		w,
		http.StatusOK,
		map[string]interface{}{"posts": posts},
	)
}
