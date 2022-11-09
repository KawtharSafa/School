package repository

import (
	"fmt"
	"gorm.io/gorm"
	"school.com/packages/internal/adapter/db/entity"
	"school.com/packages/internal/adapter/db/mapper"
	"school.com/packages/internal/domain/model"
	"school.com/packages/internal/domain/repository"
)

type ClassroomRepository struct {
	db *gorm.DB
}

func ProvideClassRepository(db *gorm.DB) repository.ClassroomRepository {
	return ClassroomRepository{db: db}
}

func (s ClassroomRepository) FindAllClassrooms(filter model.PaginationFilter) ([]*model.Classroom, error) {

	filter = model.PaginationFilter{}

	var classes []*entity.Classroom

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

	return mapper.MapToClassroomArrayModel(classes), nil
}
