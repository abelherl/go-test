package controllers

import (
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
	// Get all posts from the database
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the posts as a JSON response
	var postResponses []responses.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, responses.NewPostResponse(post))
	}
	c.JSON(200, responses.PostToJSONs(postResponses))
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
