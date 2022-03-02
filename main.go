package main

import (
	"github.com/gin-gonic/gin"

	"/controller"
	"/service"
)

var (
	studentService    service.StudentSevice        = service.New()
	studentController controller.StudentController = controller.new(studentService)
)

func main() {

	server := gin.Default()
	server.GET("/vks", func(c *gin.Context) {
		c.JSON(200, studentController.FindAll())
	})
	server.POST("/vks", func(c *gin.Context) {
		c.JSON(200, studentController.Save(c))
	})
	server.Run(":8000")
}
