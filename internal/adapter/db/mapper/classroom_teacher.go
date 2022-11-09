package mapper

import (
	"school.com/packages/internal/adapter/db/entity"
	"school.com/packages/internal/domain/model"
)

func MapToClassTeacherArrayModel(ClassroomTeacherEntity []*entity.ClassroomTeacher) []*model.ClassroomTeacher {

	var ClassroomTeacherModel []*model.ClassroomTeacher

	for _, classroomTeacher := range ClassroomTeacherEntity {

		ClassroomTeacherModel = append(ClassroomTeacherModel, &model.ClassroomTeacher{

			Teacher: model.Teacher{
				FirstName:   classroomTeacher.Teacher.FirstName,
				LastName:    classroomTeacher.Teacher.LastName,
				Description: classroomTeacher.Teacher.Description,
			},

			Classroom: model.Classroom{
				Grade: classroomTeacher.Classroom.Grade,
			},
		})
	}

	return ClassroomTeacherModel
}
