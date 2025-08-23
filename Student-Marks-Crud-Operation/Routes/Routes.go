
package Routes

import (
	"github.com/gin-gonic/gin"
	controllers "www.github.com/Aryan1365/day3/Controllers"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	studentsGroup := r.Group("/students")
	{
		studentsGroup.GET("/", controllers.GetAllStudents)
		studentsGroup.POST("/", controllers.CreateNewStudent)

		
		studentByID := studentsGroup.Group("/:id")
		{
			studentByID.GET("/", controllers.GetStudentById)
			studentByID.PUT("/", controllers.UpdateStudent)
			studentByID.DELETE("/", controllers.DeleteStudent)

			marksGroup := studentByID.Group("/marks")
			{
				marksGroup.GET("/", controllers.GetMarksByStudentId)
				marksGroup.POST("/", controllers.CreateMark)
				marksGroup.PUT("/:markId", controllers.UpdateMarks)
				marksGroup.DELETE("/:markId", controllers.DeleteMark)
			}
		}
	}
	return r
}
