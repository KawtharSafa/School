package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"school.com/packages/internal/domain/model"
	"school.com/packages/internal/interface/api/dto"
	"school.com/packages/internal/usecase/query"
)

type Classroom struct {
	GetAllClassrooms query.GetAllClassrooms
}

func ProvideClassroom(GetAllClassrooms query.GetAllClassrooms) Classroom {
	return Classroom{GetAllClassrooms: GetAllClassrooms}
}

func (s Classroom) GetClassrooms(e echo.Context) error {

	filter := model.PaginationFilter{}

	err := e.Bind(&filter)
	if err != nil {
		fmt.Println(err)
	}

	classrooms, err := s.GetAllClassrooms.Handle(filter)

	if err != nil {
		fmt.Println(err)
		return e.JSON(http.StatusInternalServerError, nil)
	}

	return e.JSON(http.StatusOK, dto.ToClassroomArrayDTO(classrooms))

}
