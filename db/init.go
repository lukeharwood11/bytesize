package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	pwd := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=5432 sslmode=disable", host, user, pwd)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
