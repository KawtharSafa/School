package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"school.com/packages/config"
	"school.com/packages/entity"
	"school.com/packages/mapper"
	"school.com/packages/model"
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

//Order("student.score desc").
//Limit(4)

//Joins("join classroom ON student.classroom_id = classroom.id").
//Group("student.id, classroom.grade").
//Having("student.score = (?)", db.
//	Table("student").
//	Joins("join classroom ON student.classroom_id = classroom.id").
//	Select("MAX(score)").
//    Where("")
//	).
//Find(&students)

//subQuery1 := db.Model(&students).Select("first_name, score")
//subQuery2 := db.Model(&classroom).Select("MAX(grade)")
//query.Table("(?) as u, (?) as p", subQuery1, subQuery2).Find(&students)

//Table("(?) as topStd", db.Model(students).
//	Select("MAX(score)")).
//Find(students)

//		Select("first_name, score, grade").
//		Joins(
//			Select(*, ROW_NUMBER() OVER (PARTITION("classroom.grade"). ORDER("student.score desc")) rn).
//			Joins("Join classroom ON (classroom.id = student.classroom_id)")
//).
//	Where(rn = 1);

//	//Statement.Build()
//
//	//subQuery := db.Select(`SELECT *, ROW_NUMBER()OVER(PARTITION BY classroom.grade ORDER BY student.score DESC as rn`).
//	//	Joins(`JOIN classroom ON classroom.id = student.classroom_id`)
//	//
//	//query.
//	//	Select("first_name, score, classroom.grade ").
//	//	Joins("join classroom ON classroom.id = student.classroom_id").
//	//	Having(subQuery)
//
//	//db.Query(`SELECT first_name, score, grade
//	//			FROM(  SELECT *, ROW_NUMBER()OVER(PARTITION BY classroom.grade ORDER BY student.score DESC) rn
//	//					FROM student INNER JOIN classroom ON (classroom.id = student.classroom_id))X
//	//		   WHERE rn = 1 `)
