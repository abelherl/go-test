package controllers

import (
	"strconv"
	"strings"

	"github.com/abelherl/go-test/initializers"
	"github.com/abelherl/go-test/models"
	"github.com/abelherl/go-test/responses"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data from request body
	var body struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	c.Bind(&body)

	// Create a new post in the database
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the created post as a JSON response
	c.JSON(200, responses.PostToJSON(responses.NewPostResponse(post)))
}

func PostsIndex(c *gin.Context) {
	// Parse query params with default values
	page, limit, search := getIndexParams(c)

	if page <= 0 || limit <= 0 {
		c.JSON(400, gin.H{"error": "Invalid pagination params"})
		return
	}

	// Calculate offset
	offset := (page - 1) * limit

	var posts []models.Post
	var total int64

	// Initialize the query
	query := initializers.DB.Model(&models.Post{})

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

	postResponses := responses.NewPostResponseList(posts)

	c.JSON(200, gin.H{
		"data":       postResponses,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": totalPages,
		"isLast":     isLast,
	})
}

func PostsShow(c *gin.Context) {
	// Get the post ID from the URL parameters
	id := c.Param("id")

	// Find the post in the database
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	// Return the post as a JSON response
	c.JSON(200, responses.PostToJSON(responses.NewPostResponse(post)))
}

func PostsUpdate(c *gin.Context) {
	// Get the post ID from the URL parameters
	id := c.Param("id")

	// Get data from request body
	var body struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	c.Bind(&body)

	// Update the post in the database
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	post.Title = body.Title
	post.Body = body.Body

	initializers.DB.Save(&post)

	// Return the updated post as a JSON response
	c.JSON(200, responses.PostToJSON(responses.NewPostResponse(post)))
}

func PostsDelete(c *gin.Context) {
	// Get the post ID from the URL parameters
	id := c.Param("id")

	// Delete the post from the database
	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	// Return a success message as a JSON response
	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}

func getIndexParams(c *gin.Context) (page int, limit int, search string) {
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
