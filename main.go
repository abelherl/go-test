package main

import (
	"github.com/abelherl/go-test/controllers"
	"github.com/abelherl/go-test/initializers"
	"github.com/abelherl/go-test/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.RateLimiter())

	// Public routes
	router.POST("/auth/login", controllers.AuthLogin)
	router.POST("/users", controllers.UserCreate)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.RequireAuth)
	{
		// Posts
		protected.POST("/posts", controllers.PostsCreate)
		protected.GET("/posts", controllers.PostsIndex)
		protected.GET("/posts/:id", controllers.PostsShow)
		protected.PUT("/posts/:id", controllers.PostsUpdate)
		protected.DELETE("/posts/:id", controllers.PostsDelete)

		// Users
		protected.GET("/users", controllers.UserIndex)
		protected.GET("/users/:id", controllers.UserShow)
		protected.PUT("/users/:id", controllers.UserUpdate)
		protected.DELETE("/users/:id", controllers.UserDelete)
	}

	router.Run()
}
