package controller

import (
	"fmt"
	"net/http"
	"school.com/packages/service/entity"
	"school.com/packages/service/mapper"
	model2 "school.com/packages/service/model"

	"github.com/labstack/echo/v4"

	"school.com/packages/config"
)

// Classrooms get classrooms info(grade)
func Classrooms(e echo.Context) error {
	classes, _ := GetClassrooms(e)
	return e.JSON(http.StatusOK, classes)
}

func GetClassrooms(e echo.Context) ([]model2.ClassroomModel, error) {

	db := config.NewDB()

	filter := model2.PaginationFilter{}

	//create classes variable
	var classes []entity.Classroom

	err := e.Bind(&filter)
	if err != nil {
		fmt.Println(err)
	}

	query := db.Table("classroom")

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

	if err := query.Find(&classes).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return mapper.MapToClassroomArrayModel(classes), nil
}
