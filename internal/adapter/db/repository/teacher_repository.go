package repository

import (
	"fmt"
	"gorm.io/gorm"
	"school.com/packages/config"
	"school.com/packages/internal/adapter/db/entity"
	"school.com/packages/internal/adapter/db/mapper"
	"school.com/packages/internal/domain/model"
	"school.com/packages/internal/domain/repository"
)

type TeacherRepository struct {
	db *gorm.DB
}

func ProvideTeacherRepository(db *gorm.DB) repository.TeacherRepository {
	return TeacherRepository{db: db}
}

func (s TeacherRepository) FindAllTeachers(filter model.PaginationFilter) ([]*model.ClassroomTeacher, error) {

	//create db new cnx
	db := config.NewDB()

	// creat pagination var
	filter = model.PaginationFilter{}

	//create variables
	var classTeacher []*entity.ClassroomTeacher

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
