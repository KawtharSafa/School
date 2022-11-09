package dto

import "school.com/packages/internal/domain/model"

type ClassroomTeacher struct {
	ClassroomId int       `gorm:"primaryKey"`
	Classroom   Classroom `gorm:"foreignKey:ClassroomId"`

	TeacherId int     `gorm:"primaryKey"`
	Teacher   Teacher `gorm:"foreignKey:TeacherId;"`
}

func ToClassroomTeacherDtO(classroomTeacher *model.ClassroomTeacher) *ClassroomTeacher {

	if nil == classroomTeacher {
		return nil
	}

	return &ClassroomTeacher{
		ClassroomId: classroomTeacher.Classroom.ClassroomId,
		TeacherId:   classroomTeacher.Teacher.TeacherId,
	}
}

func ToClassroomTeacherArrayDTO(classroomTeachers []*model.ClassroomTeacher) []*ClassroomTeacher {

	if nil == classroomTeachers {
		return nil
	}

	var classroomTeacherDTOs []*ClassroomTeacher

	for _, classroomTeacher := range classroomTeachers {
		classroomTeacherDTOs = append(classroomTeacherDTOs, ToClassroomTeacherDtO(classroomTeacher))
	}
	return classroomTeacherDTOs
}
