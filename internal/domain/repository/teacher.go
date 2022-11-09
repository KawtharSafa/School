package repository

import (
	"school.com/packages/internal/domain/model"
)

type TeacherRepository interface {
	FindAllTeachers(filter model.PaginationFilter) ([]*model.ClassroomTeacher, error)
}
