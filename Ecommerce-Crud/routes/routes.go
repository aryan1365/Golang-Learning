package routes

import (
	"retail-app/controllers"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	api := router.Group("/api")

	api.POST("/products", controllers.AddProduct)
	api.PUT("/products/:id", controllers.UpdateProduct)
	api.GET("/products", controllers.GetAllProducts)

}
