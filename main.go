package main

import (
	"github.com/Vikash-Sinha/gin_api/controller"
	"github.com/Vikash-Sinha/gin_api/service"

	"github.com/gin-gonic/gin"
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
