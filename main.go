package main

import (
	"github.com/gin-gonic/gin"
	"golang-restful-api/configs"
)
import "golang-restful-api/controllers"

func main() {
	r := gin.Default()
	configs.ConnectDatabase()

	r.GET("/ping", controllers.Pong)
	r.GET("/products", controllers.GetProducts)
	r.GET("/product/:id", controllers.GetSingleProduct)
	r.POST("/product", controllers.CreateProduct)
	r.PATCH("/product/:id", controllers.UpdateProduct)
	r.DELETE("/product/:id", controllers.DeleteProduct)

	err := r.Run()
	if err != nil {
		return
	}

}
