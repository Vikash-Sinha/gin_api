package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/Vikash-Sinha/gin_api/entity"
	"github.com/Vikash-Sinha/gin_api/service"
)

type StudentController interface{
	FindAll() []entity.Student
	Save(ctx *gin.Context) entity.Student
}
type controller struct{
	service service.StudentSevice
}
type New(service service.StudentSevice) StudentController{

	return controller{

		service:service,
	}
}

func (c *controller) FindAll() []entity.Student{
	return c.server.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entity.Student{
	var std entity.Student
	ctx.BindJSON(&std)
	c.server.Save(std)
	return std
}


type controller interface{

	
}