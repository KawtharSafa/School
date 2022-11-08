//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"school.com/packages/controller/api"
	"school.com/packages/controller/repository"
	repository2 "school.com/packages/internal/adapter/db/repository"
)

func WireStudentController(db *gorm.DB) api.StudentController {
	wire.Build(
		api.ProvideStudentController,
		repository2.ProvideStudentRepository,
	)
	return api.StudentController{}
}

func WireTeacherController(db *gorm.DB) api.TeacherController {
	wire.Build(
		api.ProvideTeacherController,
		repository.ProvideTeacherRepository,
	)
	return api.TeacherController{}
}

func WireClassroomController(db *gorm.DB) api.ClassroomController {
	wire.Build(
		api.ProvideClassroomController,
		repository.ProvideClassroomRepository,
	)
	return api.ClassroomController{}
}

func WireClassTeacherController(db *gorm.DB) api.ClassTeacherController {
	wire.Build(
		api.ProvideClassTeacherController,
		repository.ProvideClassTeacherRepository,
	)
	return api.ClassTeacherController{}
}
