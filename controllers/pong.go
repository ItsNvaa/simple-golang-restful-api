package controllers

import "github.com/gin-gonic/gin"
import "net/http"

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messege": "pong",
	})
}
