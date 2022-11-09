package mapper

import (
	"school.com/packages/internal/adapter/db/entity"
	"school.com/packages/internal/domain/model"
)

func MapToTeacherArrayModel(teacherEntity []*entity.Teacher) []*model.Teacher {

	var TeacherModel []*model.Teacher

	for _, teacher := range teacherEntity {

		TeacherModel = append(TeacherModel, &model.Teacher{
			FirstName:   teacher.FirstName,
			LastName:    teacher.LastName,
			Description: teacher.Description,
		})
	}

	return TeacherModel
}
