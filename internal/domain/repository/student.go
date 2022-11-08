package repository

import "school.com/packages/internal/domain/model"

type StudentRepository interface {
	FindAllStudents(filter model.PaginationFilter) ([]*model.Student, error)
}
