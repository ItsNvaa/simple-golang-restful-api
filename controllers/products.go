package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang-restful-api/utils"
	"gorm.io/gorm"
	"io"
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

	if isValidRequestid {
		utils.ReturnNumbericExpection(c)
		return
	}

	var product models.Product

	result := configs.DB.First(&product, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.ReturnNotFoundException(c)
			return
		}

		utils.ReturnInternalServerExpection(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
		"status":  "OK",
	})
}

func CreateProduct(c *gin.Context) {
	var productRequestBody models.Product

	err := c.ShouldBindJSON(&productRequestBody)

	if err != nil {
		if err == io.EOF {
			utils.ReturnEmptyRequestBody(c)
			return
		}

		utils.ReturnInvalidRequestBodyType(c)
		return
	}

	if err == nil {
		configs.DB.Create(&productRequestBody)
	}

	utils.ReturnOKResponse(c, "Created", http.StatusCreated)
}

func UpdateProduct(c *gin.Context) {
	var id = c.Param("id")

	var productRequestBody models.Product

	err := c.ShouldBindJSON(&productRequestBody)

	if err != nil {
		if err == io.EOF {
			utils.ReturnEmptyRequestBody(c)
			return
		}

		utils.ReturnInvalidRequestBodyType(c)
		return
	}

	var validateRequestIdResult = utils.ValidateRequestIdParams(id)

	if validateRequestIdResult {
		utils.ReturnNumbericExpection(c)
		return
	}

	var updateProduct = configs.DB.Model(&productRequestBody).Where("id = ?", id).Updates(&productRequestBody)

	if updateProduct.RowsAffected == 0 {
		utils.ReturnNotFoundException(c)
		return
	}

	utils.ReturnOKResponse(c, "Updated", http.StatusOK)
}

func DeleteProduct(c *gin.Context) {
	var id = c.Param("id")

	var product models.Product

	var validateRequestIdResult = utils.ValidateRequestIdParams(id)

	if validateRequestIdResult {
		utils.ReturnNumbericExpection(c)
		return
	}

	var deleteProduct = configs.DB.Delete(&product, id)

	if deleteProduct.RowsAffected == 0 {
		utils.ReturnNotFoundException(c)
		return
	}

	utils.ReturnOKResponse(c, "Deleted", http.StatusOK)
}
