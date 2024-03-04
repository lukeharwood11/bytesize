package models

import "gorm.io/gorm"

var ReactionType = map[string]int{
	"like":          0,
	"heart":         1,
	"support":       2,
	"informational": 3,
}

type Reaction struct {
	gorm.Model
	Article Article
	Type    int
}

type Comment struct {
	gorm.Model
	User    User
	Text    string
	Article Article
}
