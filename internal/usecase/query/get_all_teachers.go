package query

import (
	"school.com/packages/internal/domain/model"
	"school.com/packages/internal/domain/repository"
)

type GetAllTeachers struct {
	TeacherRepository repository.TeacherRepository
}

func ProvideGetAllTeachers(TeacherRepository repository.TeacherRepository) GetAllTeachers {
	return GetAllTeachers{TeacherRepository: TeacherRepository}
}

func (g GetAllTeachers) Handle(filter model.PaginationFilter) ([]*model.ClassroomTeacher, error) {
	return g.TeacherRepository.FindAllTeachers(filter)
}
