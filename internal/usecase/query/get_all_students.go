package query

import (
	"school.com/packages/internal/domain/model"
	"school.com/packages/internal/domain/repository"
)

type GetAllStudents struct {
	StudentRepository repository.StudentRepository
}

func ProvideGetAllStudents(StudentRepository repository.StudentRepository) GetAllStudents {
	return GetAllStudents{StudentRepository: StudentRepository}
}

func (g GetAllStudents) Handle(filter model.PaginationFilter) ([]*model.Student, error) {
	return g.StudentRepository.FindAllStudents(filter)
}
