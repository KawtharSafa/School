package dto

import "school.com/packages/internal/domain/model"

type Teacher struct {
	Id          int `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Description string
}

func ToTeacherDtO(teacher *model.Teacher) *Teacher {

	if nil == teacher {
		return nil
	}

	return &Teacher{
		FirstName:   teacher.FirstName,
		LastName:    teacher.LastName,
		Description: teacher.Description,
	}
}

func ToTeacherArrayDTO(teachers []*model.Teacher) []*Teacher {

	if nil == teachers {
		return nil
	}

	var teacherDTOs []*Teacher

	for _, teacher := range teachers {
		teacherDTOs = append(teacherDTOs, ToTeacherDtO(teacher))
	}
	return teacherDTOs
}
