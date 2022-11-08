package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"school.com/packages/internal/domain/model"
	"school.com/packages/internal/interface/api/dto"
	"school.com/packages/internal/usecase/query"
)

type Student struct {
	GetAllStudents query.GetAllStudents
}

func ProvideStudent(GetAllStudents query.GetAllStudents) Student {
	return Student{GetAllStudents: GetAllStudents}
}

func (s Student) GetStudents(e echo.Context) error {
	// creat pagination var
	filter := model.PaginationFilter{}

	//bind pagination filters
	err := e.Bind(&filter)
	if err != nil {
		fmt.Println(err)
	}

	//create variables
	students, err := s.GetAllStudents.Handle(filter)
	if err != nil {
		fmt.Println(err)
		return e.JSON(http.StatusInternalServerError, nil)
	}

	//return mapper.MapToStudentArrayModel(students), nil
	return e.JSON(http.StatusOK, dto.ToStudentArrayDTO(students))

}
