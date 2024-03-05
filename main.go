package main

import (
	"bytesize/controllers"
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
	env := os.Getenv("GIN_MODE")
	if err != nil && (env == "" || env != "release") {
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
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})

	api := r.Group("/api")

	{
		api.GET("/user", controllers.CreateNewUser)
	}
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
