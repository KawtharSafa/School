package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"school.com/packages/controller/repository"
	"school.com/packages/service/mapper"
	"school.com/packages/service/model"
)

type TeacherController struct {
	TeacherRepository repository.TeacherRepository
}

func (s TeacherController) GetAllTeachers(e echo.Context) error {

	filter := model.PaginationFilter{}

	err := e.Bind(&filter)
	if err != nil {
		fmt.Println(err)
	}

	var AllTeachers = s.TeacherRepository.FindTeachers()

	return e.JSON(http.StatusOK, mapper.MapToTeacherArrayModel(AllTeachers))
}

func ProvideTeacherController(TeacherRepository repository.TeacherRepository) TeacherController {
	return TeacherController{TeacherRepository: TeacherRepository}
}
