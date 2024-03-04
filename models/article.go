package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string
	Subtitle string
	Author   User
	Url      string
	Content  string
}
