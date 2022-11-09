package mapper

import (
	"school.com/packages/internal/adapter/db/entity"
	"school.com/packages/internal/domain/model"
)

func MapToClassroomArrayModel(classroomEntity []*entity.Classroom) []*model.Classroom {

	var ClassroomModel []*model.Classroom

	for _, classroom := range classroomEntity {

		ClassroomModel = append(ClassroomModel, &model.Classroom{
			Grade: classroom.Grade,
		})
	}

	return ClassroomModel
}
