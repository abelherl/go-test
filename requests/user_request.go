package requests

import (
	"github.com/abelherl/go-test/helpers"
	"github.com/abelherl/go-test/models"
)

type UserRequest struct {
	FirstName string `json:"firstName" gorm:"column:first_name"`
	LastName  string `json:"lastName"  gorm:"column:last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (r UserRequest) ToUserModel() (models.User, error) {
	password, err := helpers.HashPassword(r.Password)

	return models.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Password:  password,
	}, err
}
