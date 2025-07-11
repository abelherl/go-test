package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
