package models

// defines the permissions of each user
var Role = map[string]int{
	"admin":   0,
	"creator": 1,
	"default": 2,
}

type User struct {
	BaseModel
	Name     string
	Password string
	Username string
	Role     int `gorm:"default:2"`
}
