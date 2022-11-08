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

type ClassTeacherController struct {
	ClassTeacherRepository repository.ClassTeacherRepository
}

func (s ClassTeacherController) GetClassTeacher(e echo.Context) error {

	filter := model.PaginationFilter{}

	err := e.Bind(&filter)
	if err != nil {
		fmt.Println(err)
	}

	var ClassTeacher = s.ClassTeacherRepository.FindClassTeacher()

	return e.JSON(http.StatusOK, mapper.MapToClassTeacherArrayModel(ClassTeacher))
}

func ProvideClassTeacherController(ClassTeacherRepository repository.ClassTeacherRepository) ClassTeacherController {
	return ClassTeacherController{ClassTeacherRepository: ClassTeacherRepository}
}
