package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
}
