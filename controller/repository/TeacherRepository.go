package repository

import (
	"gorm.io/gorm"
	"school.com/packages/service/model"
)

type TeacherRepository struct {
	db *gorm.DB
}

func (s TeacherRepository) FindTeachers() []*model.TeacherModel {
	//should return array of teachers as json teacher model
	return nil
}

func ProvideTeacherRepository(db *gorm.DB) TeacherRepository {
	return TeacherRepository{db: db}
}
