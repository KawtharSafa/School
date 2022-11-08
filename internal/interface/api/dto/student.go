package dto

import "school.com/packages/internal/domain/model"

type Student struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func ToStudentDtO(student *model.Student) *Student {
	if nil == student {
		return nil
	}
	return &Student{
		FirstName: student.FirstName,
		LastName:  student.LastName,
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
