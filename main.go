package main

import (
	"fmt"
	"go-go/app/controller"
	"go-go/app/route"
	"go-go/lib"
	"net/http"
	"os"
)

func main() {
	lib.LoadEnv()

	port := os.Getenv("PORT") 
	if port == "" {
		port = "8080"
	}
	
	routes := route.NewRouter()
	routes.GET("/", controller.Index)
	routes.GET("/ping", controller.Ping)


	fmt.Println("Server running on port", port)
	http.ListenAndServe(
		fmt.Sprintf(":%s", port),
		routes,
	)

}
