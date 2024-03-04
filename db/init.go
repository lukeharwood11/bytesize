package db

import (
	"bytesize/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	// grab environment variables
	user := os.Getenv("POSTGRES_USER")
	pwd := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	// create connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=5432 sslmode=disable", host, user, pwd)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// this will also call os.Exit(1)
		log.Fatal("Could not connect to the database.")
	}

	DB.AutoMigrate(&models.User{}, &models.Article{}, &models.Reaction{}, &models.Comment{})
}
