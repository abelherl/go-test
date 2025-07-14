package controllers

import (
	"github.com/abelherl/go-test/initializers"
	"github.com/abelherl/go-test/models"
	"github.com/abelherl/go-test/requests"
	"github.com/abelherl/go-test/responses"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func UserCreate(c *gin.Context) {
	// Get data from request body
	var body requests.UserRequest

	c.Bind(&body)

	// Prepare data to model
	user, err := body.ToUserModel()

	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create a new user in the database
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		if pgErr, ok := result.Error.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case "23505": // unique_violation
				c.JSON(400, gin.H{"error": "Email already exists"})
				return
			}
		}

		// Fallback: generic DB error
		c.JSON(400, gin.H{"error": "Failed to create user"})
		return
	}

	// Return the created user as a JSON response
	c.JSON(200, responses.UserToJSON(user))
}

func UserIndex(c *gin.Context) {
	// Get all users from the database
	var users []models.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Convert users to response format
	userResponses := responses.UserToJSONList(users)

	// Return the list of users as a JSON response
	c.JSON(200, userResponses)
}

func UserShow(c *gin.Context) {
	// Get user ID from URL
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	// Find user by ID in the database
	var user models.User
	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Return the user as a JSON response
	c.JSON(200, responses.UserToJSON(user))
}

func UserUpdate(c *gin.Context) {
	// Get user ID from URL
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	// Get data from request body
	var body requests.UserRequest

	c.Bind(&body)

	// Prepare data to model
	user, err := body.ToUserModel()
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to hash password"})
		return
	}

	// Update user in the database
	result := initializers.DB.Model(&models.User{}).Where("id = ?", id).Updates(user)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to update user"})
		return
	}

	// Return the updated user as a JSON response
	c.JSON(200, responses.UserToJSON(user))
}

func UserDelete(c *gin.Context) {
	// Get user ID from URL
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	// Delete user from the database
	result := initializers.DB.Delete(&models.User{}, id)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to delete user"})
		return
	}

	// Return a success message as a JSON response
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
