package repository

import (
	"school.com/packages/internal/domain/model"
)

type ClassroomRepository interface {
	FindAllClassrooms(filter model.PaginationFilter) ([]*model.Classroom, error)
}
