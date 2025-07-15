package main

import (
	"github.com/abelherl/go-test/controllers"
	"github.com/abelherl/go-test/initializers"
	"github.com/abelherl/go-test/middleware"
	"github.com/abelherl/go-test/services"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.InitCloudinary()
}

func main() {
	// Initialize dependencies
	db := initializers.DB

	userService := services.NewUserService(db)

	authController := controllers.NewAuthController(userService)
	postController := controllers.NewPostController(db)
	userController := controllers.NewUserController(db)

	// Set up Gin router
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.RateLimiter())

	// Public routes
	router.POST("/auth/login", authController.AuthLogin)
	router.POST("/users", userController.UserCreate)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.RequireAuth)
	{
		// Posts
		protected.POST("/posts", postController.PostsCreate)
		protected.GET("/posts", postController.PostsIndex)
		protected.GET("/posts/:id", postController.PostsShow)
		protected.PUT("/posts/:id", postController.PostsUpdate)
		protected.DELETE("/posts/:id", postController.PostsDelete)

		// Users
		protected.GET("/users", userController.UserIndex)
		protected.GET("/users/:id", userController.UserShow)
		protected.PUT("/users/:id", userController.UserUpdate)
		protected.DELETE("/users/:id", userController.UserDelete)
	}

	router.Run()
}
