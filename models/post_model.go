package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string
	Body        string
	Attachments pq.StringArray `gorm:"type:text[]" json:"attachments"`
}
