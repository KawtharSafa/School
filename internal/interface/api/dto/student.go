package dto

import "school.com/packages/internal/domain/model"

type Student struct {
	Id          int `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Score       float32
	ClassroomId int
	Classroom   Classroom `gorm:"foreignKey:ClassroomId"`
	Grade       string
}

func ToStudentDtO(student *model.Student) *Student {

	if nil == student {
		return nil
	}

	return &Student{
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Score:     student.Score,
		Grade:     student.Classroom.Grade,
	}
}

func ToStudentArrayDTO(students []*model.Student) []*Student {

	if nil == students {
		return nil
	}

	var studentDTOs []*Student

	for _, student := range students {
		studentDTOs = append(studentDTOs, ToStudentDtO(student))
	}
	return studentDTOs
}
