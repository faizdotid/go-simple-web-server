package route

import (
	"fmt"
	"go-go/app/helper"
	"log"
	"net/http"
	"strings"
)

// Router struct to manage route
type router struct {
	rt       []route
	notfound map[string]interface{}
}

// Create new router
func NewRouter() *router {
	return &router{
		rt:       []route{},
		notfound: map[string]interface{}{"message": "Not Found"},
	}
}

// not found handler
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

func (r *router) setResponseHeaderAsJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func (r *router) LogRequest(req *http.Request) {
	log.Printf("IP: %s, Method: %s, Path: %s", req.RemoteAddr, req.Method, req.URL.Path)
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// log request into console
	r.LogRequest(req)

	// logic for handling route
	request_path_explode := strings.Split(req.URL.Path, "/")
	for _, route := range r.rt {
		route_path_explode := strings.Split(route.path, "/")

		if len(route_path_explode) != len(request_path_explode) {
			continue
		}
		// checking method first
		if route.method != req.Method {
			helper.JSON(
				w,
				http.StatusMethodNotAllowed,
				map[string]interface{}{"message": fmt.Sprintf("Method %s not allowed", req.Method)},
			)
			return
		}

		route.handler(w, req)
		return
	}

	// route not found
	helper.JSON(w, http.StatusNotFound, r.notfound)
}
