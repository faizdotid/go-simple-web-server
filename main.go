package main

import (
	"fmt"
	"go-go/app/controller"
	"go-go/app/middleware"
	"go-go/app/model"
	"go-go/app/route"
	"go-go/app/service"
	_ "go-go/lib"
	"log"
	"net/http"
	"os"
)

func main() {
	// recover from panic
	// and exit the program
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()

	// create new database service
	db := service.NewDatabaseService()
	db.Connect()

	// defer close db connection
	defer db.Close()

	// create new router
	routes := route.NewRouter()

	// create new user model
	userModel := model.NewUserModel(db.GetDB())
	postModel := model.NewPostModel(db.GetDB())

	//route index
	routes.GET("/", controller.Index)
	routes.GET("/ping", controller.Ping)

	// auth routes
	{
		authController := controller.NewAuthController(userModel)
		routes.POST("/login", authController.Login)       // done
		routes.POST("/register", authController.Register) // done
	}

	// posts routes
	{
		postController := controller.NewPostController(postModel)
		routes.GET("/posts", postController.Index)
		routes.GET("/posts/search/:query", postController.Search)
		routes.GET("/post/:id", postController.Show) //
		routes.POST("/post", middleware.Authenticate(postController.Create))
		routes.PUT("/post/:id", func(w http.ResponseWriter, r *http.Request) {})
		routes.DELETE("/post/:id", func(w http.ResponseWriter, r *http.Request) {})
	}

	// user routes
	{
		userController := controller.NewUserController(userModel)
		routes.GET("/me", middleware.Authorize(userController.Index))
		routes.GET("/me/posts", func(w http.ResponseWriter, r *http.Request) {})
	}

	// run the server
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", routes)
}
