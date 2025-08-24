package controllers

import (
	"net/http"

	"git.com/crud/initializers"
	"git.com/crud/models"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}
func PostsCreate(c *gin.Context) {
	var pst models.Post
	c.BindJSON(&pst)
	// post := models.Post{
	// 	Name: "First Post",
	// 	Body: "This is the body of the first post",
	// }
	result := initializers.DB.Create(&pst)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": pst,
	})

}
