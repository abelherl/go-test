package services

import (
	"github.com/abelherl/go-test/initializers"
	"github.com/abelherl/go-test/models"
)

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := initializers.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}
