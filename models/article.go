package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Title    string
	Subtitle string
	AuthorID uuid.UUID
	Author   User
	Url      string
	Content  string
}
