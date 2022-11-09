package query

import (
	"school.com/packages/internal/domain/model"
	"school.com/packages/internal/domain/repository"
)

type GetAllClassrooms struct {
	ClassroomRepository repository.ClassroomRepository
}

func ProvideGetAllClassrooms(ClassroomRepository repository.ClassroomRepository) GetAllClassrooms {
	return GetAllClassrooms{ClassroomRepository: ClassroomRepository}
}

func (g GetAllClassrooms) Handle(filter model.PaginationFilter) ([]*model.Classroom, error) {
	return g.ClassroomRepository.FindAllClassrooms(filter)
}
