// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"gorm.io/gorm"
	"school.com/packages/internal/adapter/db/repository"
	"school.com/packages/internal/interface/api"
	"school.com/packages/internal/usecase/query"
)

// Injectors from wire.go:

func WireStudentController(db *gorm.DB) api.Student {
	studentRepository := repository.ProvideStudentRepository(db)
	getAllStudents := query.ProvideGetAllStudents(studentRepository)
	student := api.ProvideStudent(getAllStudents)
	return student
}

func WireTeacherController(db *gorm.DB) api.Teacher {
	teacherRepository := repository.ProvideTeacherRepository(db)
	getAllTeachers := query.ProvideGetAllTeachers(teacherRepository)
	teacher := api.ProvideTeacher(getAllTeachers)
	return teacher
}

func WireClassroomController(db *gorm.DB) api.Classroom {
	classroomRepository := repository.ProvideClassRepository(db)
	getAllClassrooms := query.ProvideGetAllClassrooms(classroomRepository)
	classroom := api.ProvideClassroom(getAllClassrooms)
	return classroom
}

func WireClassTeacherController(db *gorm.DB) api.ClassroomTeacher {
	classroomTeacherRepository := repository.ProvideClassroomTeacherRepository(db)
	getAllClassroomTeacher := query.ProvideGetAllClassroomTeacher(classroomTeacherRepository)
	classroomTeacher := api.ProvideClassroomTeacher(getAllClassroomTeacher)
	return classroomTeacher
}
