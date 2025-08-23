
package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"www.github.com/Aryan1365/day3/Config"
	models "www.github.com/Aryan1365/day3/Models"
)

func GetAllStudents(c *gin.Context) {
	var students []models.Student
	err := models.GetStudents(&students)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch students"})
		return
	}
	c.JSON(200, students)
}

func CreateNewStudent(c *gin.Context) {
	var student models.Student
	c.BindJSON(&student)
	err := models.CreateStudent(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusCreated, student)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	err := models.GetUserById(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&student)
	err = models.UpdateStudent(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	err := models.DeleteStudent(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})

}

func GetStudentById(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	err := models.GetUserById(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, student)
}

func CreateMark(c *gin.Context) {
	studentId := c.Param("id")
	var marks models.Mark
	id, err := strconv.ParseUint(studentId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid studentId"})
		return
	}
	marks.StudentId = uint(id)
	if err := c.ShouldBindJSON(&marks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// c.BindJSON(&marks)

	if err := models.CreateMark(&marks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create mark"})
		return
	}
	fmt.Println(marks)
	var student models.Student
	if err := Config.DB.Preload("Marks").First(&student, uint(id)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch student"})
		return
	}
	c.JSON(http.StatusCreated, student)
}

func GetMarksByStudentId(c *gin.Context) {
	studentId := c.Param("id")
	var student models.Student
	err := models.GetUserById(&student, studentId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, student.Marks)
}

func UpdateMarks(c *gin.Context) {
	markId := c.Param("markId")
	var mark models.Mark
	err := models.GetMarksInfo(&mark, markId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&mark)
	err = models.UpdateMarks(&mark)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, mark)

}

func DeleteMark(c *gin.Context) {
	markId := c.Param("markId")
	err := models.DeleteMark(markId)
	if err != nil {
		if errors.Is(err, errors.New("mark not found")) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Mark not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete mark"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mark deleted successfully"})
}
