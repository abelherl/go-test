package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/abelherl/go-test/helpers"
	"github.com/abelherl/go-test/initializers"
	"github.com/abelherl/go-test/models"
	"github.com/abelherl/go-test/requests"
	"github.com/abelherl/go-test/responses"
	"github.com/abelherl/go-test/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	DB      *gorm.DB
	Service *services.UserService
}

func NewPostController(db *gorm.DB, service *services.UserService) *PostController {
	return &PostController{DB: db, Service: service}
}

func (pc PostController) PostsCreate(c *gin.Context) {
	var body requests.PostsCreate

	c.Bind(&body)

	userID, err := helpers.GetUserIDFromAuth(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to get user ID"})
		return
	}

	post, err := requests.NewPostFromCreateRequest(body, userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to parse request"})
		return
	}

	result := pc.DB.Create(&post)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(200, responses.PostToJSON(responses.NewPostResponse(pc.Service, post)))
}

func (pc PostController) PostsIndex(c *gin.Context) {
	// Parse query params with default values
	page, limit, search := pc.getIndexParams(c)

	if page <= 0 || limit <= 0 {
		c.JSON(400, gin.H{"error": "Invalid pagination params"})
		return
	}

	// Calculate offset
	offset := (page - 1) * limit

	var posts []models.Post
	var total int64

	// Initialize the query
	query := pc.DB.Model(&models.Post{}).Where("author_id IS NOT NULL")

	// Apply search filter if provided
	if search != "" {
		query = query.Where("COALESCE(title, '') ILIKE ? OR COALESCE(body, '') ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Count total filtered records
	query.Count(&total)
	totalPages := (total + int64(limit) - 1) / int64(limit)
	isLast := page >= int(totalPages)

	// Fetch paginated filtered results
	query.
		Limit(limit).
		Offset(offset).
		Find(&posts)

	postResponses := responses.NewPostResponseList(pc.Service, posts)

	c.JSON(200, gin.H{
		"data":       postResponses,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": totalPages,
		"isLast":     isLast,
	})
}

func (pc PostController) PostsShow(c *gin.Context) {
	// Get the post ID from the URL parameters
	id := c.Param("id")

	// Find the post in the database
	var post models.Post
	result := pc.DB.First(&post, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	// Return the post as a JSON response
	c.JSON(200, responses.PostToJSON(responses.NewPostResponse(pc.Service, post)))
}

func (pc PostController) PostsUpdate(c *gin.Context) {
	id := c.Param("id")
	var body requests.PostsUpdate

	c.Bind(&body)

	var post models.Post
	result := pc.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(401, gin.H{"error": "Failed to find post"})
		return
	}

	isSameAuthor := checkSameAuthor(c, post.AuthorID)
	if !isSameAuthor {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No access to this resource"})
		return
	}

	newPost, err := requests.NewPostFromUpdateRequest(body, id, post.AuthorID)

	if err != nil {
		if err.Error() == "invalid tags or technologies" {
			c.JSON(400, gin.H{"error": "Invalid tags or techs"})
			return
		}

		c.JSON(400, gin.H{"error": "Failed to parse request"})
		return
	}

	pc.DB.Save(&newPost)

	var updatedPost models.Post
	pc.DB.First(&updatedPost, id)

	c.JSON(200, responses.PostToJSON(responses.NewPostResponse(pc.Service, updatedPost)))
}

func (pc PostController) PostsDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	resultFirst := pc.DB.First(&post, id)
	if resultFirst.Error != nil {
		c.JSON(401, gin.H{"error": "Failed to find post"})
		return
	}

	isSameAuthor := checkSameAuthor(c, post.AuthorID)
	if !isSameAuthor {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No access to this resource"})
		return
	}

	result := pc.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}

func (pc PostController) PostsUploadAttachments(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	result := pc.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to get post data"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to parse form data"})
		return
	}

	files := form.File["attachments"]
	if len(files) == 0 {
		c.JSON(400, gin.H{"error": "No attachments provided"})
		return
	} else if len(files) > 5 {
		c.JSON(400, gin.H{"error": "Maximum 5 attachments allowed"})
		return
	}

	for i, file := range files {
		index := strconv.Itoa(i)
		fileName := "post_" + id + "_" + index
		url, err := initializers.UploadImage(c.Request.Context(), file, fileName, "post_attachments")
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to upload attachment"})
			return
		}
		post.Attachments = append(post.Attachments, url)
	}

	pc.DB.Save(&post)

	c.JSON(200, responses.PostToJSON(responses.NewPostResponse(pc.Service, post)))
}

func (pc PostController) getIndexParams(c *gin.Context) (page int, limit int, search string) {
	page = 1
	limit = 20
	search = strings.TrimSpace(c.Query("search"))

	if p, err := strconv.Atoi(c.Query("page")); err == nil && p > 0 {
		page = p
	}

	if l, err := strconv.Atoi(c.Query("limit")); err == nil && l > 0 {
		limit = l
	}

	return
}

func checkSameAuthor(c *gin.Context, authorID uint) bool {
	userID, err := helpers.GetUserIDFromAuth(c)
	if err != nil {
		return false
	}
	return userID == authorID
}
