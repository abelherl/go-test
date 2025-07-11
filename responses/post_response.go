package responses

import (
	"github.com/abelherl/go-test/models"
	"github.com/gin-gonic/gin"
)

type PostResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewPostResponse(p models.Post) PostResponse {
	return PostResponse{
		ID:    p.ID,
		Title: p.Title,
		Body:  p.Body,
	}
}

func PostToJSON(post PostResponse) gin.H {
	return gin.H{
		"data": post,
	}
}

func PostToJSONs(posts []PostResponse) gin.H {
	return gin.H{
		"data": posts,
	}
}
