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
	postController := controllers.NewPostController(db, userService)
	userController := controllers.NewUserController(db)

	// Set up Gin router
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.RateLimiter())

	// Public routes
	router.POST("/auth/login", authController.AuthLogin)
	router.POST("/users", userController.UserCreate)
	router.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	// Routes with general authentication check
	authGroup := router.Group("/")
	authGroup.Use(middleware.RequireAuth)

	authGroup.POST("/posts", postController.PostsCreate)
	authGroup.GET("/posts", postController.PostsIndex)
	authGroup.GET("/posts/:id", postController.PostsShow)
	authGroup.PUT("/posts/:id", postController.PostsUpdate)
	authGroup.PUT("/posts/:id/upload-attachments", postController.PostsUploadAttachments)
	authGroup.DELETE("/posts/:id", postController.PostsDelete)

	authGroup.GET("/users", userController.UserIndex)
	authGroup.GET("/users/:id", userController.UserShow)

	// Routes with same user authentication check
	userGroup := router.Group("/")
	userGroup.Use(middleware.RequireAuthSameUser)

	userGroup.PUT("/users/:id", userController.UserUpdate)
	userGroup.PUT("/users/:id/upload-profile-photo", userController.UserUploadProfilePhoto)
	userGroup.DELETE("/users/:id", userController.UserDelete)

	router.Run()
}
