package mapper

import (
	"school.com/packages/service/entity"
	"school.com/packages/service/model"
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

func MapToClassroomArrayModel(classroomEntity []*entity.Classroom) []*model.Classroom {

	var ClassroomModel []*model.Classroom

	for _, classroom := range classroomEntity {

		ClassroomModel = append(ClassroomModel, &model.Classroom{
			Grade: classroom.Grade,
		})
	}

	return ClassroomModel
}

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

func MapToStudentArrayModel(studentEntity []*entity.Student) []*model.Student {

	var StudentModel []*model.Student

	for _, student := range studentEntity {

		StudentModel = append(StudentModel, &model.Student{
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Score:     student.Score,

			Classroom: model.Classroom{
				Grade: student.Classroom.Grade,
			},
		})
	}

	return StudentModel
}
