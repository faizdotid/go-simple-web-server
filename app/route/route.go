package route

import (
	"net/http"
	// "regexp"
	// "strings"
)

// Route struct to define the route
type route struct {
	path    string
	method  string
	handler http.HandlerFunc
}

// func (r *route) Params(path_params, path_routes []string) map[string]string {
// 	params := map[string]string{}
// 	for i, route := range path_routes {
// 		if value, _ := regexp.MatchString("{.*}", route); value {
// 			params[route[1:len(route)-1]] = path_params[i]
// 		}
// 	}
// 	return map[string]string{}
// }
