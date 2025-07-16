package requests

import (
	"errors"
	"strconv"

	"github.com/abelherl/go-test/helpers"
	"github.com/abelherl/go-test/models"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type PostsCreate struct {
	Body         string         `json:"body"`
	Title        string         `json:"title"`
	Tags         pq.StringArray `json:"tags"`
	TechFrontEnd pq.StringArray `json:"techFrontEnd"`
	TechBackEnd  pq.StringArray `json:"techBackEnd"`
	TechInfra    pq.StringArray `json:"techInfra"`
	TechNextGen  pq.StringArray `json:"techNextGen"`
}

type PostsUpdate struct {
	Body         string         `json:"body"`
	Title        string         `json:"title"`
	Tags         pq.StringArray `json:"tags"`
	TechFrontEnd pq.StringArray `json:"techFrontEnd"`
	TechBackEnd  pq.StringArray `json:"techBackEnd"`
	TechInfra    pq.StringArray `json:"techInfra"`
	TechNextGen  pq.StringArray `json:"techNextGen"`
}

func NewPostFromCreateRequest(request PostsCreate) (models.Post, error) {
	if !helpers.ValidateTags(request.Tags) ||
		!helpers.ValidateTechnologies(request.TechFrontEnd) ||
		!helpers.ValidateTechnologies(request.TechBackEnd) ||
		!helpers.ValidateTechnologies(request.TechInfra) ||
		!helpers.ValidateTechnologies(request.TechNextGen) {
		return models.Post{}, errors.New("invalid tags or technologies")
	}

	return models.Post{
		Title:        request.Title,
		Body:         request.Body,
		Tags:         request.Tags,
		TechFrontEnd: request.TechFrontEnd,
		TechBackEnd:  request.TechBackEnd,
		TechInfra:    request.TechInfra,
		TechNextGen:  request.TechNextGen,
	}, nil
}

func NewPostFromUpdateRequest(request PostsUpdate, id string) (models.Post, error) {
	if !helpers.ValidateTags(request.Tags) ||
		!helpers.ValidateTechnologies(request.TechFrontEnd) ||
		!helpers.ValidateTechnologies(request.TechBackEnd) ||
		!helpers.ValidateTechnologies(request.TechInfra) ||
		!helpers.ValidateTechnologies(request.TechNextGen) {
		return models.Post{}, errors.New("invalid tags or technologies")
	}

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return models.Post{}, errors.New("invalid post ID")
	}

	return models.Post{
		Model: gorm.Model{
			ID: uint(parsedID),
		},
		Title:        request.Title,
		Body:         request.Body,
		Tags:         request.Tags,
		TechFrontEnd: request.TechFrontEnd,
		TechBackEnd:  request.TechBackEnd,
		TechInfra:    request.TechInfra,
		TechNextGen:  request.TechNextGen,
	}, nil
}
