//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"school.com/internal/adapter/repository"
	"school.com/internal/interface/api"
	"school.com/internal/usecase/query"
)

func WireStudentController(db *gorm.DB) api.StudentController {
	wire.Build(
		api.ProvideStudent,
		query.ProvideGetAllStudents,
		repository.ProvideStudentRepository,
	)
	return api.StudentController{}
}

func WireTeacherController(db *gorm.DB) api.TeacherController {
	wire.Build(
		api.ProvideTeacher,
		query.ProvideGetAllTeachers,
		repository.ProvideTeacherRepository,
	)
	return api.TeacherController{}
}

func WireClassroomController(db *gorm.DB) api.ClassroomController {
	wire.Build(
		api.ProvideClassroom,
		query.ProvideGetAllClassrooms,
		repository.ProvideClassroomRepository,
	)
	return api.ClassroomController{}
}

func WireClassTeacherController(db *gorm.DB) api.ClassTeacherController {
	wire.Build(
		api.ProvideClassTeacher,
		query.ProvideGetAllClassroomTeacher,
		repository.ProvideClassTeacherRepository,
	)
	return api.ClassTeacherController{}
}
