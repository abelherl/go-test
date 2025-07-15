package services

import (
	"github.com/abelherl/go-test/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (us *UserService) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := us.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}
