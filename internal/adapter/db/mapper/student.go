package mapper

import (
	"school.com/packages/internal/adapter/db/entity"
	"school.com/packages/internal/domain/model"
)

func MapToStudentArrayModel(studentEntity []*entity.Student) []*model.Student {

	if nil == studentEntity {
		return nil
	}

	var StudentModel []*model.Student

	for _, student := range studentEntity {

		StudentModel = append(StudentModel, &model.Student{
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Score:     student.Score,

			Classroom: model.ClassroomModel{
				Grade: student.Classroom.Grade,
			},
		})
	}

	return StudentModel
}
