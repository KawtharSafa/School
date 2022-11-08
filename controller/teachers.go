package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"school.com/packages/config"
	"school.com/packages/service/entity"
	"school.com/packages/service/mapper"
	model2 "school.com/packages/service/model"
)

// Teachers find the teachers in class x
func Teachers(e echo.Context) error {
	teachersData, _ := GetTeachers(e)
	return e.JSON(http.StatusOK, teachersData)
}

func GetTeachers(e echo.Context) ([]model2.ClassroomTeacherModel, error) {

	//create db new cnx
	db := config.NewDB()

	// creat pagination var
	filter := model2.PaginationFilter{}

	//create variables
	var classTeacher []entity.ClassroomTeacher

	//bind pagination filters
	err := e.Bind(&filter)
	if err != nil {
		fmt.Println(err)
	}

	//get teachers of such class
	query := db.
		Joins("join classroom ON classroom.id = classroom_teacher.classroom_id").
		Joins("join teacher ON teacher.id = classroom_teacher.teacher_id").
		Preload("Classroom").
		Preload("Teacher")

	//joins of tables if filter is by grade
	if filter.Classroom != "" {

		query.
			//Select("teacher.first_name, teacher.last_name, teacher.description, classroom.grade").
			Where("classroom.grade = ?", filter.Classroom)

	}

	//tables if filter is order by teacher's name in 2 directions
	if filter.SortBy == "Teachers" {
		if filter.Direction == "ASC" {
			query.
				Order("teacher.first_name ASC")
		}
		if filter.Direction == "DSC" {
			query.
				Order("teacher.first_name desc")
		}
	}

	//tables if filter is order by grade in 2 directions
	if filter.SortBy == "Classes" {
		if filter.Direction == "ASC" {
			query.
				Order("classroom.grade ASC")
		}
		if filter.Direction == "DSC" {
			query.
				Order("classroom.grade desc")
		}
	}

	if err := query.
		Find(&classTeacher).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return mapper.MapToClassTeacherArrayModel(classTeacher), nil
}
