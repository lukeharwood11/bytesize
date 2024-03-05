package models

import uuid "github.com/satori/go.uuid"

var ReactionType = map[string]int{
	"like":          0,
	"heart":         1,
	"support":       2,
	"informational": 3,
}

type Reaction struct {
	BaseModel
	Article   Article
	ArticleID uuid.UUID
	Type      int
}

type Comment struct {
	BaseModel
	User      User
	UserID    uuid.UUID
	Text      string
	Article   Article
	ArticleID uint
}
