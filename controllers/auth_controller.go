package controllers

import (
	"net/http"

	"github.com/abelherl/go-test/helpers"
	"github.com/abelherl/go-test/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService services.UserServiceInterface
}

func NewAuthController(userService services.UserServiceInterface) *AuthController {
	return &AuthController{UserService: userService}
}

func (ac *AuthController) AuthLogin(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := ac.UserService.GetUserByEmail(body.Email)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !helpers.CheckPassword(body.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
