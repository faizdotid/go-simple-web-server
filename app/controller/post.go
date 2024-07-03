package controller

import (
	"encoding/json"
	"errors"
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
	Create(user_id int, post model.Post) (model.Post, error)
	Delete(id int) error
	Update(id int, post model.Post) (model.Post, error)
}

type postControllers struct {
	postModel postModelInterface
}

func NewPostController(postModel postModelInterface) *postControllers {
	return &postControllers{postModel: postModel}
}

func (a *postControllers) RequiredForm(r *http.Request, required ...string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, errors.New("invalid request body")
	}
	if len(data) == 0 {
		return nil, errors.New("request body is empty")
	}
	for k, v := range data {
		if ok := helper.InArray(k, required); ok && v == nil {
			return nil, errors.New("all fields are required")
		}
	}
	return data, nil
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
	user := r.Context().Value("user")
	var id int
	if user == nil {
		id = 0
	} else {
		id = user.(map[string]interface{})["id"].(int)
	}
	data, err := p.RequiredForm(r, "title", "content")
	if err != nil {
		helper.JSON(
			w,
			http.StatusBadRequest,
			map[string]interface{}{"error": err.Error()},
		)
		return
	}
	post := model.Post{
		Title:   data["title"].(string),
		Content: data["content"].(string),
	}
	post, err = p.postModel.Create(id, post)
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
		http.StatusCreated,
		map[string]interface{}{"post": post},
	)
}

func (p *postControllers) Delete(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	if user == nil {
		helper.JSON(w, http.StatusUnauthorized, map[string]interface{}{"error": "you are not authorized to delete this post"})
		return
	}
	params := r.Context().Value(route.ParamsKey("params")).(route.Params)
	id := params.Get("id").(int)
	helper.JSON(w, http.StatusOK, map[string]interface{}{"id": id})

}

func (p *postControllers) Update(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	if user == nil {
		helper.JSON(w, http.StatusUnauthorized, map[string]interface{}{"error": "you are not authorized to update this post"})
		return
	}
	params := r.Context().Value(route.ParamsKey("params")).(route.Params)
	id := params.Get("id").(int)
	data, err := p.RequiredForm(r, "title", "content")
	if err != nil {
		helper.JSON(
			w,
			http.StatusBadRequest,
			map[string]interface{}{"error": err.Error()},
		)
		return
	}
	post := model.Post{
		Title:   data["title"].(string),
		Content: data["content"].(string),
	}
	post, err = p.postModel.Update(id, post)
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
		map[string]interface{}{"post": post},
	)
}