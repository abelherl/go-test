package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title        string
	Body         string
	Tags         pq.StringArray `json:"tags" gorm:"type:text[]"`
	TechFrontEnd pq.StringArray `json:"techFrontEnd" gorm:"type:text[]"`
	TechBackEnd  pq.StringArray `json:"techBackEnd" gorm:"type:text[]"`
	TechInfra    pq.StringArray `json:"techInfra" gorm:"type:text[]"`
	TechNextGen  pq.StringArray `json:"techNextGen" gorm:"type:text[]"`
	Attachments  pq.StringArray `gorm:"type:text[]" json:"attachments"`
}
