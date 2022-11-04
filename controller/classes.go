package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"school.com/packages/config"
	"school.com/packages/entity"
	"school.com/packages/mapper"
	"school.com/packages/model"
)

// Classrooms get classrooms info(grade)
func Classrooms(e echo.Context) error {
	classes, _ := GetClassrooms(e)
	return e.JSON(http.StatusOK, classes)
}

func GetClassrooms(e echo.Context) ([]model.ClassroomModel, error) {

	db := config.NewDB()

	filter := model.PaginationFilter{}

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
