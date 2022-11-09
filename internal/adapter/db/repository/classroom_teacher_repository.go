package repository

import (
	"fmt"
	"gorm.io/gorm"
	"school.com/packages/internal/adapter/db/entity"
	"school.com/packages/internal/adapter/db/mapper"
	"school.com/packages/internal/domain/model"
	"school.com/packages/internal/domain/repository"
)

type ClassroomTeacherRepository struct {
	db *gorm.DB
}

func ProvideClassroomTeacherRepository(db *gorm.DB) repository.ClassroomTeacherRepository {
	return ClassroomTeacherRepository{db: db}
}

func (s ClassroomTeacherRepository) FindAllClassroomTeacher(filter model.PaginationFilter) ([]*model.ClassroomTeacher, error) {

	filter = model.PaginationFilter{}

	var classes []*entity.ClassroomTeacher

	query := s.db.Table("classroom")

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

	return mapper.MapToClassTeacherArrayModel(classes), nil
}
