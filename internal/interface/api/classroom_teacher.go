package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"school.com/packages/internal/domain/model"
	"school.com/packages/internal/interface/api/dto"
	"school.com/packages/internal/usecase/query"
)

type ClassroomTeacher struct {
	GetAllClassroomTeacher query.GetAllClassroomTeacher
}

func ProvideClassroomTeacher(GetAllClassroomTeacher query.GetAllClassroomTeacher) ClassroomTeacher {
	return ClassroomTeacher{GetAllClassroomTeacher: GetAllClassroomTeacher}
}

func (s ClassroomTeacher) GetClassroomTeacher(e echo.Context) error {

	filter := model.PaginationFilter{}

	err := e.Bind(&filter)
	if err != nil {
		fmt.Println(err)
	}

	classroomTeacher, err := s.GetAllClassroomTeacher.Handle(filter)

	if err != nil {
		fmt.Println(err)
		return e.JSON(http.StatusInternalServerError, nil)
	}

	return e.JSON(http.StatusOK, dto.ToClassroomTeacherArrayDTO(classroomTeacher))

}
