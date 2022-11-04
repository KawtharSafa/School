package mapper

import (
	"school.com/packages/entity"
	"school.com/packages/model"
)

func MapToClassTeacherArrayModel(ClassroomTeacherEntity []entity.ClassroomTeacher) []model.ClassroomTeacherModel {

	var ClassroomTeacherModel []model.ClassroomTeacherModel

	for _, classroomTeacher := range ClassroomTeacherEntity {

		ClassroomTeacherModel = append(ClassroomTeacherModel, model.ClassroomTeacherModel{

			Teacher: model.TeacherModel{
				FirstName:   classroomTeacher.Teacher.FirstName,
				LastName:    classroomTeacher.Teacher.LastName,
				Description: classroomTeacher.Teacher.Description,
			},

			Classroom: model.ClassroomModel{
				Grade: classroomTeacher.Classroom.Grade,
			},
		})
	}

	return ClassroomTeacherModel
}

func MapToClassroomArrayModel(classroomEntity []entity.Classroom) []model.ClassroomModel {

	var ClassroomModel []model.ClassroomModel

	for _, classroom := range classroomEntity {

		ClassroomModel = append(ClassroomModel, model.ClassroomModel{
			Grade: classroom.Grade,
		})
	}

	return ClassroomModel
}

func MapToTeacherArrayModel(teacherEntity []entity.Teacher) []model.TeacherModel {

	var TeacherModel []model.TeacherModel

	for _, teacher := range teacherEntity {

		TeacherModel = append(TeacherModel, model.TeacherModel{
			FirstName:   teacher.FirstName,
			LastName:    teacher.LastName,
			Description: teacher.Description,
		})
	}

	return TeacherModel
}

func MapToStudentArrayModel(studentEntity []entity.Student) []model.StudentModel {

	var StudentModel []model.StudentModel

	for _, student := range studentEntity {

		StudentModel = append(StudentModel, model.StudentModel{
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
