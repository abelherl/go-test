package middleware

import (
	"strconv"

	"github.com/abelherl/go-test/helpers"
	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	tokenStr := helpers.GetAuthHeader(c)
	if tokenStr == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	token, err := helpers.ValidateJWT(tokenStr)
	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}

func RequireAuthSameUser(c *gin.Context) {
	userID, err := helpers.GetUserIDFromAuth(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	requestedUserID := c.Param("id")
	if requestedUserID == "" || requestedUserID != strconv.Itoa(int(userID)) {
		c.JSON(403, gin.H{"error": "No access"})
		c.Abort()
		return
	}

	c.Next()
}
