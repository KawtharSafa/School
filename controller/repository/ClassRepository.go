package repository

import (
	"gorm.io/gorm"
	"school.com/packages/service/model"
)

type ClassroomRepository struct {
	db *gorm.DB
}

func (s ClassroomRepository) FindClasses() []*model.ClassroomModel {
	//should return array of classrooms as json classroom model
	return nil
}

func ProvideClassroomRepository(db *gorm.DB) ClassroomRepository {
	return ClassroomRepository{db: db}
}
