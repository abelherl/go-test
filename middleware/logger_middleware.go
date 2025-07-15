package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Process request
		c.Next()

		// Log after request
		duration := time.Since(start)
		status := c.Writer.Status()

		log.Printf("[GIN] %d | %v | %s %s", status, duration, method, path)
	}
}
