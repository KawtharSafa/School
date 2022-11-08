package repository

import (
	"gorm.io/gorm"
	"school.com/packages/service/model"
)

type ClassTeacherRepository struct {
	db *gorm.DB
}

func (s ClassTeacherRepository) FindClassTeacher() []*model.ClassTeacherModel {
	//should return array of ClassTeacher as json ClassTeacher model
	return nil
}

func ProvideClassTeacherRepository(db *gorm.DB) ClassTeacherRepository {
	return ClassTeacherRepository{db: db}
}
