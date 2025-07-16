package responses

import (
	"github.com/abelherl/go-test/helpers"
	"github.com/abelherl/go-test/models"
	"github.com/abelherl/go-test/services"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type PostResponse struct {
	ID           uint           `json:"id"`
	Title        string         `json:"title"`
	Body         string         `json:"body"`
	Author       AuthorResponse `json:"author"`
	Tags         pq.StringArray `json:"tags" gorm:"type:text[]"`
	TechFrontEnd pq.StringArray `json:"techFrontEnd" gorm:"type:text[]"`
	TechBackEnd  pq.StringArray `json:"techBackEnd" gorm:"type:text[]"`
	TechInfra    pq.StringArray `json:"techInfra" gorm:"type:text[]"`
	TechNextGen  pq.StringArray `json:"techNextGen" gorm:"type:text[]"`
	Attachments  pq.StringArray `json:"attachments" gorm:"type:text[]" `
}

type AuthorResponse struct {
	ID              uint   `json:"id"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	ProfilePhotoURL string `json:"profilePhotoURL"`
}

func NewAuthorResponse(id uint, firstName, lastName, photoURL string) AuthorResponse {
	return AuthorResponse{
		ID:              id,
		FirstName:       firstName,
		LastName:        lastName,
		ProfilePhotoURL: photoURL,
	}
}

func NewPostResponse(s *services.UserService, p models.Post) PostResponse {
	user := s.GetUserByID(p.AuthorID)
	author := NewAuthorResponse(user.ID, user.FirstName, user.LastName, user.ProfilePhotoURL)

	return PostResponse{
		ID:           p.ID,
		Title:        p.Title,
		Body:         p.Body,
		Author:       author,
		Tags:         helpers.EmptyPQStringArrayIfNil(p.Tags),
		TechFrontEnd: helpers.EmptyPQStringArrayIfNil(p.TechFrontEnd),
		TechBackEnd:  helpers.EmptyPQStringArrayIfNil(p.TechBackEnd),
		TechInfra:    helpers.EmptyPQStringArrayIfNil(p.TechInfra),
		TechNextGen:  helpers.EmptyPQStringArrayIfNil(p.TechNextGen),
		Attachments:  helpers.EmptyPQStringArrayIfNil(p.Attachments),
	}
}

func NewPostResponseList(service *services.UserService, posts []models.Post) []PostResponse {
	responses := make([]PostResponse, 0, len(posts))
	for _, post := range posts {
		responses = append(responses, NewPostResponse(service, post))
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
