package repository

import (
	"gorm.io/gorm"
	"school.com/packages/service/model"
)

type StudentRepository struct {
	db *gorm.DB
}

func (s StudentRepository) FindAllStudents() []*model.StudentModel {
	//should return array of students as json model
	return nil
}

func ProvideStudentRepository(db *gorm.DB) StudentRepository {
	return StudentRepository{db: db}
}
