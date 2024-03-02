package main

import (
	"bytesize/db"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	env := os.Getenv("APP_ENV")
	if err != nil && (env == "" || env == "dev") {
		log.Fatal("Error loading .env file")
	}

	// initialize db connection
	db, err := db.InitDatabase()
	fmt.Println(db.Name())
	if err != nil {
		log.Fatal("Could not connect to the database")
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}