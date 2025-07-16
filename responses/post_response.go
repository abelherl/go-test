package responses

import (
	"github.com/abelherl/go-test/helpers"
	"github.com/abelherl/go-test/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type PostResponse struct {
	ID           uint           `json:"id"`
	Title        string         `json:"title"`
	Body         string         `json:"body"`
	Tags         pq.StringArray `json:"tags" gorm:"type:text[]"`
	TechFrontEnd pq.StringArray `json:"techFrontEnd" gorm:"type:text[]"`
	TechBackEnd  pq.StringArray `json:"techBackEnd" gorm:"type:text[]"`
	TechInfra    pq.StringArray `json:"techInfra" gorm:"type:text[]"`
	TechNextGen  pq.StringArray `json:"techNextGen" gorm:"type:text[]"`
	Attachments  pq.StringArray `json:"attachments" gorm:"type:text[]" `
}

func NewPostResponse(p models.Post) PostResponse {
	return PostResponse{
		ID:           p.ID,
		Title:        p.Title,
		Body:         p.Body,
		Tags:         helpers.EmptyPQStringArrayIfNil(p.Tags),
		TechFrontEnd: helpers.EmptyPQStringArrayIfNil(p.TechFrontEnd),
		TechBackEnd:  helpers.EmptyPQStringArrayIfNil(p.TechBackEnd),
		TechInfra:    helpers.EmptyPQStringArrayIfNil(p.TechInfra),
		TechNextGen:  helpers.EmptyPQStringArrayIfNil(p.TechNextGen),
		Attachments:  helpers.EmptyPQStringArrayIfNil(p.Attachments),
	}
}

func NewPostResponseList(posts []models.Post) []PostResponse {
	responses := make([]PostResponse, 0, len(posts))
	for _, post := range posts {
		responses = append(responses, NewPostResponse(post))
	}
	return responses
}

func PostToJSON(post PostResponse) gin.H {
	return gin.H{
		"data": post,
	}
}

func PostToJSONList(posts []PostResponse) gin.H {
	return gin.H{
		"data": posts,
	}
}
