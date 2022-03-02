package service

import "/entity"

type StudentSevice interface {
	Save(entity.Student) entity.Student
	FindAll() []entity.Student
}

type studentService struct {
	students []entity.Student
}

func New() StudentSevice {

	return &studentService{}
}

func (service *studentService) Save(student entity.Student) entity.Student {
	service.students = append(service.students, student)
	return student
}
func (service *studentService) FindAll() []entity.Student {
	return service.students
}
