package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnNumbericExpection(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"TypeError": "Bad Request",
		"messege":   "The Request Id Must Be Numberic",
	})
}

func ReturnInternalServerExpection(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"TypeError": "Internal Server Error",
	})
}

func ReturnOKResponse(c *gin.Context, field string, status int) {
	c.JSON(status, gin.H{
		field:    true,
		"status": "OK",
	})
}

func ReturnEmptyRequestBody(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"TypeError": "Bad Request",
		"messege":   "The Request Body Should Be Defined",
	})
}

func ReturnInvalidRequestBodyType(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"TypeError": "Bad Request",
		"messege":   "The Request Body is Invalid!",
	})
}

func ReturnNotFoundException(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"TypeError": "Not Found",
		"messege":   "The Data Was Not Found",
	})
}
