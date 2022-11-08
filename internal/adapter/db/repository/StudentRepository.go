package repository

import (
	"fmt"
	"gorm.io/gorm"
	"school.com/packages/internal/adapter/db/entity"
	"school.com/packages/internal/adapter/db/mapper"
	"school.com/packages/internal/domain/model"
	"school.com/packages/internal/domain/repository"
)

type StudentRepository struct {
	db *gorm.DB
}

func ProvideStudentRepository(db *gorm.DB) repository.StudentRepository {
	return StudentRepository{db: db}
}

func (s StudentRepository) FindAllStudents(filter model.PaginationFilter) ([]*model.Student, error) {
	var students []*entity.Student
	//get page nb
	page := filter.Page
	offset := (page - 1) * 5

	//get students of such class
	query := s.db.Preload("Classroom")

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
