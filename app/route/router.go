package route

import (
	"context"
	"fmt"
	"go-go/app/helper"
	"log"
	"net/http"
	"strings"
)

type Params map[string]interface{}

type ParamsKey string

// Get value from Params
func (p Params) Get(key string) interface{} {
	if val, ok := p[key]; ok {
		return val
	}
	return ""
}

// router struct to manage route
type router struct {
	rt       []route
	notfound map[string]interface{}
}

// create new router
func NewRouter() *router {
	return &router{
		rt:       []route{},
		notfound: map[string]interface{}{"message": "Not Found"},
	}
}

// handle not found route
func (r *router) NotFound(data map[string]interface{}) {
	r.notfound = data
}

func (r *router) GET(path string, handler http.HandlerFunc) {
	r.rt = append(r.rt, route{path: path, method: "GET", handler: handler})
}

func (r *router) POST(path string, handler http.HandlerFunc) {
	r.rt = append(r.rt, route{path: path, method: "POST", handler: handler})
}

func (r *router) PUT(path string, handler http.HandlerFunc) {
	r.rt = append(r.rt, route{path: path, method: "PUT", handler: handler})
}

func (r *router) DELETE(path string, handler http.HandlerFunc) {
	r.rt = append(r.rt, route{path: path, method: "DELETE", handler: handler})
}

func (r *router) LogRequest(req *http.Request) {
	log.Printf("IP: %s, Method: %s, Path: %s", req.RemoteAddr, req.Method, req.URL.Path)
}

// ServeHTTP to serve http request
func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// log request into console
	r.LogRequest(req)

	// trim trailing slash
	req.URL.Path = strings.TrimSuffix(req.URL.Path, "/")
	request_path_explode := strings.Split(req.URL.Path, "/")

	// loop through registered routes
	for _, route := range r.rt {

		// trim trailing slash
		route.path = strings.TrimSuffix(route.path, "/")
		route_path_explode := strings.Split(route.path, "/")

		// get allowed method in same route
		allowed_method_in_same_route := []string{}
		for _, method := range r.rt {
			if strings.Contains(method.path, route.path) {
				allowed_method_in_same_route = append(allowed_method_in_same_route, method.method)
			}
		}

		// check if route path and request path has same length
		if len(route_path_explode) != len(request_path_explode) {
			continue
		}

		// check if route path and request path has same value
		match := true
		params := Params{}
		paramsKey := ParamsKey("params")
		for i, route_path := range route_path_explode {
			if strings.HasPrefix(route_path, ":") {
				params[route_path[1:]] = request_path_explode[i]
				continue
			}
			if route_path != request_path_explode[i] {
				match = false
				break
			}
		}

		// if not match, continue to next route
		if !match {
			continue
		}

		// checking method allowed in same route
		if !helper.InArray(req.Method, allowed_method_in_same_route) {
			helper.JSON(
				w,
				http.StatusMethodNotAllowed,
				map[string]interface{}{"message": fmt.Sprintf("Method %s not allowed", req.Method)},
			)
			return
		}

		// checking method first
		if req.Method != route.method {
			continue
		}

		// create new context with params
		ctx := context.WithValue(req.Context(), ParamsKey(paramsKey), params)
		route.handler(w, req.WithContext(ctx))
		return
	}

	// handle not found route
	helper.JSON(
		w,
		http.StatusNotFound,
		r.notfound,
	)
}
