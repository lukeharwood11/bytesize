package models

import "gorm.io/gorm"

// defines the permissions of each user
var Role = map[string]int{
	"admin":   0,
	"creator": 1,
	"default": 2,
}

type User struct {
	gorm.Model
	Name     string
	Password string
	Username string
	Role     int `gorm:"default:2"`
}
