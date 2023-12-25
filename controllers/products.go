package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang-restful-api/utils"
	"gorm.io/gorm"
)
import "net/http"
import "golang-restful-api/models"
import "golang-restful-api/configs"

func GetProducts(c *gin.Context) {
	var products []models.Product

	configs.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"status":   "OK",
	})
}

func GetSingleProduct(c *gin.Context) {
	var id = c.Param("id")

	var isValidRequestid = utils.ValidateRequestIdParams(id)

	if !isValidRequestid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"TypeError": "Bad Request",
			"messege":   "The Request Id Must Be Numberic",
		})
		return
	}

	var product models.Product

	result := configs.DB.First(&product, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"TypeError": "Not Found",
				"messege":   "The Data Wasn't Found",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"TypeError": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
		"status":  "OK",
	})
}
