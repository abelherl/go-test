package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abelherl/go-test/controllers"
	"github.com/abelherl/go-test/helpers"
	"github.com/abelherl/go-test/models"
	"github.com/abelherl/go-test/tests/mocks/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthLoginSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUserService := new(services.MockUserService)
	authController := controllers.NewAuthController(mockUserService)

	// Password hash for "password123"
	hashedPassword, _ := helpers.HashPassword("password123")

	mockUser := &models.User{
		Email:    "test@example.com",
		Password: hashedPassword,
	}

	mockUserService.
		On("GetUserByEmail", "test@example.com").
		Return(mockUser, nil)

	body := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}
	bodyJSON, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	authController.AuthLogin(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}

func TestAuthLoginInvalidPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUserService := new(services.MockUserService)
	authController := controllers.NewAuthController(mockUserService)

	hashedPassword, _ := helpers.HashPassword("correctPassword")
	mockUser := &models.User{
		Email:    "test@example.com",
		Password: hashedPassword,
	}

	mockUserService.
		On("GetUserByEmail", "test@example.com").
		Return(mockUser, nil)

	body := map[string]string{
		"email":    "test@example.com",
		"password": "wrongPassword",
	}
	bodyJSON, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	authController.AuthLogin(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
}

func TestAuthLoginUserNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUserService := new(services.MockUserService)
	authController := controllers.NewAuthController(mockUserService)

	mockUserService.
		On("GetUserByEmail", "notfound@example.com").
		Return(nil, errors.New("user not found"))

	body := map[string]string{
		"email":    "notfound@example.com",
		"password": "irrelevant",
	}
	bodyJSON, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	authController.AuthLogin(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
}
