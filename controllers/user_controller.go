package controllers

import (
	"github.com/abelherl/go-test/initializers"
	"github.com/abelherl/go-test/models"
	"github.com/abelherl/go-test/requests"
	"github.com/abelherl/go-test/responses"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) UserCreate(c *gin.Context) {
	var body requests.UserRequest
	c.Bind(&body)

	user, err := body.ToUserModel()
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to hash password"})
		return
	}

	result := uc.DB.Create(&user)
	if result.Error != nil {
		if pgErr, ok := result.Error.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			c.JSON(400, gin.H{"error": "Email already exists"})
			return
		}
		c.JSON(400, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(200, responses.UserToJSON(user))
}

func (uc *UserController) UserIndex(c *gin.Context) {
	var users []models.User
	result := uc.DB.Find(&users)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch users"})
		return
	}

	userResponses := responses.UserToJSONList(users)
	c.JSON(200, userResponses)
}

func (uc *UserController) UserShow(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	result := uc.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, responses.UserToJSON(user))
}

func (uc *UserController) UserUpdate(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	var body requests.UserRequest
	c.Bind(&body)

	user, err := body.ToUserModel()
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to hash password"})
		return
	}

	result := uc.DB.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(200, responses.UserToJSON(user))
}

func (uc *UserController) UserDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	result := uc.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func (uc *UserController) UserUploadProfilePhoto(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to get file"})
		return
	}

	ctx := c.Request.Context()
	publicID := "user_" + id

	url, err := initializers.UploadImage(ctx, file, publicID, "profile_photos")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to upload image"})
		return
	}

	var user models.User
	result := uc.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	user.ProfilePhotoURL = url
	uc.DB.Save(&user)

	c.JSON(200, responses.UserToJSON(user))
}
