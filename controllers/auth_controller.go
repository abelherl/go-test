package controllers

import (
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
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := ac.UserService.GetUserByEmail(body.Email)
	passErr := !helpers.CheckPassword(body.Password, user.Password)

	if err != nil || passErr {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
