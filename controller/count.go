package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"school.com/packages/config"
	"school.com/packages/entity"
	"school.com/packages/model"
)

func Count(e echo.Context) error {
	countData, _ := GetCount(e)
	return e.JSON(http.StatusOK, countData)
}

func GetCount(e echo.Context) (*model.ClassroomCounter, error) {

	//create db new cnx
	db := config.NewDB()

	//create variables
	var class entity.Classroom

	//filter and its bind with error
	filter := model.CountFilter{}

	err := e.Bind(&filter)
	if err != nil {
		fmt.Println(err)
	}

	response := model.ClassroomCounter{
		Classroom: filter.Classroom,
	}

	query := db.Model(&class).Select(`id, (SELECT COUNT(teacher_id) FROM classroom_teacher 
							WHERE classroom.id = classroom_teacher.classroom_id) AS teacher_count,
							(SELECT COUNT(student.id) FROM student
							WHERE classroom.id = student.classroom_id) AS student_count`).
		Group("id")

	//joins of tables if filter is by grade
	if filter.Classroom != "" {

		query.
			Where("classroom.grade = ?", filter.Classroom)
	}

	if err := query.
		First(&response).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &response, nil
}
