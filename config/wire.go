//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"school.com/packages/internal/adapter/db/repository"
	"school.com/packages/internal/interface/api"
	"school.com/packages/internal/usecase/query"
)

func WireStudentController(db *gorm.DB) api.Student {
	wire.Build(
		api.ProvideStudent,
		query.ProvideGetAllStudents,
		repository.ProvideStudentRepository,
	)
	return api.Student{}
}

func WireTeacherController(db *gorm.DB) api.Teacher {
	wire.Build(
		api.ProvideTeacher,
		query.ProvideGetAllTeachers,
		repository.ProvideTeacherRepository,
	)
	return api.Teacher{}
}

func WireClassroomController(db *gorm.DB) api.Classroom {
	wire.Build(
		api.ProvideClassroom,
		query.ProvideGetAllClassrooms,
		repository.ProvideClassRepository,
	)
	return api.Classroom{}
}

func WireClassTeacherController(db *gorm.DB) api.ClassroomTeacher {
	wire.Build(
		api.ProvideClassroomTeacher,
		query.ProvideGetAllClassroomTeacher,
		repository.ProvideClassroomTeacherRepository,
	)
	return api.ClassroomTeacher{}
}
