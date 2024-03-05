package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewUser(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"message": "hello world!",
	})
}
