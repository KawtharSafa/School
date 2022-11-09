package query

import (
	"school.com/packages/internal/domain/model"
	"school.com/packages/internal/domain/repository"
)

type GetAllClassroomTeacher struct {
	ClassroomTeacherRepository repository.ClassroomTeacherRepository
}

func ProvideGetAllClassroomTeacher(ClassroomTeacherRepository repository.ClassroomTeacherRepository) GetAllClassroomTeacher {
	return GetAllClassroomTeacher{ClassroomTeacherRepository: ClassroomTeacherRepository}
}

func (g GetAllClassroomTeacher) Handle(filter model.PaginationFilter) ([]*model.ClassroomTeacher, error) {
	return g.ClassroomTeacherRepository.FindAllClassroomTeacher(filter)
}
