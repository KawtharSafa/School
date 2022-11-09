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

	filter := model.PaginationFilter{}

	err := e.Bind(&filter)
	if err != nil {
		fmt.Println(err)
	}

	students, err := s.GetAllStudents.Handle(filter)

	//var studentsDto []*dto.Student
	//for _, student := range students {
	//	studentsDto = append(studentsDto, dto.ToStudentDtO(student))
	//}

	if err != nil {
		fmt.Println(err)
		return e.JSON(http.StatusInternalServerError, nil)
	}

	return e.JSON(http.StatusOK, dto.ToStudentArrayDTO(students))

}
