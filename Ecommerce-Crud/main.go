package main

import (
	"retail-app/config"
	"retail-app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
	routes.setUp(r)
	r.Run(":8080")
}
