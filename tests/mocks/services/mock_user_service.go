package services

import (
	"github.com/abelherl/go-test/models"
	"github.com/abelherl/go-test/services"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

var _ services.UserServiceInterface = &MockUserService{}

func (m *MockUserService) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)

	user, _ := args.Get(0).(*models.User)
	return user, args.Error(1)
}

func (m *MockUserService) GetUserByID(id uint) *models.User {
	args := m.Called(id)

	user, _ := args.Get(0).(*models.User)
	return user
}
