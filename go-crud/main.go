package main

import (
	"git.com/crud/controllers"
	"git.com/crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}
func main() {
	router := gin.Default()

	router.POST("/", controllers.PostsCreate)
	router.GET("/all", controllers.GetPost)

	router.Run() // listen and serve on 0.0.0.0:8080
}
