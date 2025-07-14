package responses

import (
	"github.com/abelherl/go-test/models"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName" gorm:"column:first_name"`
	LastName  string `json:"lastName"  gorm:"column:last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

func NewUserResponse(u models.User) UserResponse {
	if u.Role == "" {
		u.Role = "user"
	}
	return UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Role:      u.Role,
	}
}

func UserToJSON(u models.User) gin.H {
	return gin.H{
		"data": NewUserResponse(u),
	}
}

func UserToJSONList(users []models.User) gin.H {
	responses := make([]UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, NewUserResponse(user))
	}
	return gin.H{
		"data": responses,
	}
}
