package repository

import (
	"school.com/packages/internal/domain/model"
)

type ClassroomTeacherRepository interface {
	FindAllClassroomTeacher(filter model.PaginationFilter) ([]*model.ClassroomTeacher, error)
}
