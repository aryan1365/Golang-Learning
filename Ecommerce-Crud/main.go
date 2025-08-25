package main

import (
	"retail-app/config"
	"retail-app/models"
	"retail-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
	routes.Setup(r)
	r.Run(":8080")
}
