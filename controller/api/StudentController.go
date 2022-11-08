package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"school.com/packages/controller/repository"
	//"school.com/packages/service/entity"
	"school.com/packages/service/mapper"
	"school.com/packages/service/model"
)

type StudentController struct {
	StudentRepository repository.StudentRepository
}

func (s StudentController) GetAllStudents(e echo.Context) error {

	filter := model.PaginationFilter{}

	//var students []entity.Student

	err := e.Bind(&filter)
	if err != nil {
		fmt.Println(err)
	}

	var Allstudents = s.StudentRepository.FindAllStudents()

	return e.JSON(http.StatusOK, mapper.MapToStudentArrayModel(Allstudents))
}

func ProvideStudentController(StudentRepository repository.StudentRepository) StudentController {
	return StudentController{StudentRepository: StudentRepository}
}
