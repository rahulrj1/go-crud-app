package main

import (
	"go-crud-app/handlers"
	"go-crud-app/initializers"
	"go-crud-app/repository"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.LoadDatabase()
}

func main() {
	router := gin.Default()

	greetRouter := router.Group("/greet")
	greetRouter.GET("/", handlers.Greet)
	greetRouter.POST("/", handlers.GreetToName)

	postRepo := repository.NewMongoPostRepository(initializers.PostCollection)
	postHandler := handlers.NewPostHandler(postRepo)

	postRouter := router.Group("/posts")
	postRouter.GET("/", postHandler.GetAllPosts)
	postRouter.POST("/", postHandler.CreatePost)
	postRouter.GET("/:id", postHandler.GetPost)
	postRouter.PUT("/:id", postHandler.UpdatePost)
	postRouter.DELETE("/:id", postHandler.DeletePost)

	router.Run()
}
