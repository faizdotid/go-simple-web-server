package controller

import (
	"fmt"
	"go-go/app/helper"
	"go-go/app/model"
	"go-go/app/route"
	"log"
	"net/http"
)

type postModelInterface interface {
	FindAll() ([]model.Post, error)
	FindById(id int) (model.Post, error)
	Search(query string) ([]model.Post, error)
}

type postControllers struct {
	postModel postModelInterface
}

func NewPostController(postModel postModelInterface) *postControllers {
	return &postControllers{postModel: postModel}
}

func (p *postControllers) Index(w http.ResponseWriter, r *http.Request) {
	posts, err := p.postModel.FindAll()
	log.Println(posts)
	if err != nil {
		helper.JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	helper.JSON(w, http.StatusOK, map[string]interface{}{"posts": posts})
}

func (p *postControllers) Show(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(route.ParamsKey("params")).(route.Params)
	fmt.Println(params.Get("id"))
	post, err := p.postModel.FindById(1)
	if err != nil {
		helper.JSON(w, http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}
	helper.JSON(w, http.StatusOK, map[string]interface{}{"post": post})
}

func (p *postControllers) Search(w http.ResponseWriter, r *http.Request) {
	query := r.Context().Value(route.ParamsKey("params")).(route.Params).Get("query").(string)
	posts, err := p.postModel.Search(query)
	if err != nil {
		helper.JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	helper.JSON(w, http.StatusOK, map[string]interface{}{"posts": posts})
}

func (p *postControllers) Create(w http.ResponseWriter, r *http.Request) {
	helper.JSON(w, http.StatusOK, map[string]interface{}{"message": "Create"})
}