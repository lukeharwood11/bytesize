package main

import (
	"bytesize/db"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	initializeEnvironmentVariables()
	db.InitDatabase()
}

func initializeEnvironmentVariables() {
	err := godotenv.Load()
	env := os.Getenv("APP_ENV")
	if err != nil && (env == "" || env == "dev") {
		log.Fatal("Error loading .env file.")
	}
}

func main() {
	err := godotenv.Load()
	env := os.Getenv("APP_ENV")
	if err != nil && (env == "" || env == "dev") {
		log.Fatal("Error loading .env file")
	}

	if err != nil {
		log.Fatal("Could not connect to the database")
	}
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
