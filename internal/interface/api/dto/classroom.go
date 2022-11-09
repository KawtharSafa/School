package dto

import "school.com/packages/internal/domain/model"

type Classroom struct {
	Id    int    `gorm:"primaryKey"`
	Grade string `gorm:"uniqueIndex:Grade"`
}

func ToClassroomDtO(classroom *model.Classroom) *Classroom {

	if nil == classroom {
		return nil
	}

	return &Classroom{
		Grade: classroom.Grade,
	}
}

func ToClassroomArrayDTO(classrooms []*model.Classroom) []*Classroom {

	if nil == classrooms {
		return nil
	}

	var classroomDTOs []*Classroom

	for _, classroom := range classrooms {
		classroomDTOs = append(classroomDTOs, ToClassroomDtO(classroom))
	}
	return classroomDTOs
}
