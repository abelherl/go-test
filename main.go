package main

import (
	"github.com/abelherl/go-test/controllers"
	"github.com/abelherl/go-test/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	router := gin.Default()

	router.POST("/posts", controllers.PostsCreate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.PUT("/posts/:id", controllers.PostsUpdate)
	router.DELETE("/posts/:id", controllers.PostsDelete)

	router.POST("/users", controllers.UserCreate)
	router.GET("/users", controllers.UserIndex)
	router.GET("/users/:id", controllers.UserShow)
	router.PUT("/users/:id", controllers.UserUpdate)
	router.DELETE("/users/:id", controllers.UserDelete)

	router.Run()
}
