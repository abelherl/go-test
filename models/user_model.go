package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName       string
	LastName        string
	ProfilePhotoURL string
	Email           string `gorm:"uniqueIndex"`
	Password        string `json:"-"`
	Role            string `gorm:"default:'user'"`
}
