package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"school.com/packages/internal/domain/model"
	"school.com/packages/internal/interface/api/dto"
	"school.com/packages/internal/usecase/query"
)

type Teacher struct {
	GetAllTeachers query.GetAllTeachers
}

func ProvideTeacher(GetAllTeachers query.GetAllTeachers) Teacher {
	return Teacher{GetAllTeachers: GetAllTeachers}
}

func (s Teacher) GetTeachers(e echo.Context) error {

	filter := model.PaginationFilter{}

	err := e.Bind(&filter)
	if err != nil {
		fmt.Println(err)
	}

	teachers, err := s.GetAllTeachers.Handle(filter)

	if err != nil {
		fmt.Println(err)
		return e.JSON(http.StatusInternalServerError, nil)
	}

	return e.JSON(http.StatusOK, dto.ToClassroomTeacherArrayDTO(teachers))

}
