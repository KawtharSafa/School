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

type ClassroomController struct {
	ClassroomRepository repository.ClassroomRepository
}

func (s ClassroomController) GetClassrooms(e echo.Context) error {

	filter := model.PaginationFilter{}

	err := e.Bind(&filter)
	if err != nil {
		fmt.Println(err)
	}

	var AllClassrooms = s.ClassroomRepository.FindClasses()

	return e.JSON(http.StatusOK, mapper.MapToClassroomArrayModel(AllClassrooms))
}

func ProvideClassroomController(ClassroomRepository repository.ClassroomRepository) ClassroomController {
	return ClassroomController{ClassroomRepository: ClassroomRepository}
}
