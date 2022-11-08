package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"school.com/packages/config"
	"school.com/packages/service/entity"
	"school.com/packages/service/mapper"
	"school.com/packages/service/model"
)

// Students find the students in class x
func Students(e echo.Context) error {
	studentsData, _ := GetStudents(e)
	return e.JSON(http.StatusOK, studentsData)
}

func GetStudents(e echo.Context) ([]model.StudentModel, error) {

	//create db new cnx
	db := config.NewDB()

	// creat pagination var
	filter := model.PaginationFilter{}

	//create variables
	var students []entity.Student

	//bind pagination filters
	err := e.Bind(&filter)
	if err != nil {
		fmt.Println(err)
	}

	//get page nb
	page := filter.Page
	offset := (page - 1) * 5

	//get students of such class
	query := db.Preload("Classroom")

	//joins of tables if filter is by grade
	if filter.Classroom != "" {

		query.
			Joins("join classroom ON student.classroom_id = classroom.id").
			Where("classroom.grade = ?", filter.Classroom)
	}

	//tables if filter is order by student name in 2 directions
	if filter.SortBy == "Students" {
		if filter.Direction == "ASC" {
			query.
				Order("student.first_name ASC")
		}
		if filter.Direction == "DSC" {
			query.
				Order("student.first_name desc")
		}
	}

	//tables if filter is order by grade in 2 directions
	if filter.SortBy == "Classes" {
		if filter.Direction == "ASC" {
			query.
				Joins("Join classroom ON student.classroom_id = classroom.id").
				Order("classroom.grade ASC")
		}
		if filter.Direction == "DSC" {
			query.
				Joins("join classroom ON student.classroom_id = classroom.id").
				Order("classroom.grade desc")
		}
	}

	//tables if filter is top student in classes
	if filter.TopStudent == true {
		query.
			Preload("Classroom").
			Raw(`SELECT first_name, score, grade 
					FROM ( SELECT *, ROW_NUMBER()OVER(PARTITION BY classroom.grade ORDER BY student.score DESC) rn 
					       FROM student INNER JOIN classroom ON classroom.id = student.classroom_id)X 
					WHERE rn =1;`)

	}

	if filter.PerPage == true {
		query.
			Limit(5).
			Offset(offset)

	}

	if err := query.
		Find(&students).
		Error; err != nil {
		fmt.Println(err)
		return nil, err

	}

	return mapper.MapToStudentArrayModel(students), nil
}
