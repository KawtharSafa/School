package mapper

import (
	entity2 "school.com/packages/service/entity"
	model2 "school.com/packages/service/model"
)

func MapToClassTeacherArrayModel(ClassroomTeacherEntity []entity2.ClassroomTeacher) []model2.ClassroomTeacherModel {

	var ClassroomTeacherModel []model2.ClassroomTeacherModel

	for _, classroomTeacher := range ClassroomTeacherEntity {

		ClassroomTeacherModel = append(ClassroomTeacherModel, model2.ClassroomTeacherModel{

			Teacher: model2.TeacherModel{
				FirstName:   classroomTeacher.Teacher.FirstName,
				LastName:    classroomTeacher.Teacher.LastName,
				Description: classroomTeacher.Teacher.Description,
			},

			Classroom: model2.ClassroomModel{
				Grade: classroomTeacher.Classroom.Grade,
			},
		})
	}

	return ClassroomTeacherModel
}

func MapToClassroomArrayModel(classroomEntity []entity2.Classroom) []model2.ClassroomModel {

	var ClassroomModel []model2.ClassroomModel

	for _, classroom := range classroomEntity {

		ClassroomModel = append(ClassroomModel, model2.ClassroomModel{
			Grade: classroom.Grade,
		})
	}

	return ClassroomModel
}

func MapToTeacherArrayModel(teacherEntity []entity2.Teacher) []model2.TeacherModel {

	var TeacherModel []model2.TeacherModel

	for _, teacher := range teacherEntity {

		TeacherModel = append(TeacherModel, model2.TeacherModel{
			FirstName:   teacher.FirstName,
			LastName:    teacher.LastName,
			Description: teacher.Description,
		})
	}

	return TeacherModel
}

func MapToStudentArrayModel(studentEntity []entity2.Student) []model2.StudentModel {

	var StudentModel []model2.StudentModel

	for _, student := range studentEntity {

		StudentModel = append(StudentModel, model2.StudentModel{
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Score:     student.Score,

			Classroom: model2.ClassroomModel{
				Grade: student.Classroom.Grade,
			},
		})
	}

	return StudentModel
}
