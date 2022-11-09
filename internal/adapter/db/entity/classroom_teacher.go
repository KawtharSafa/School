package entity

import (
	"fmt"
)

type ClassroomTeacher struct {
	ClassroomId int       `gorm:"primaryKey"`
	Classroom   Classroom `gorm:"foreignKey:ClassroomId"`

	TeacherId int     `gorm:"primaryKey"`
	Teacher   Teacher `gorm:"foreignKey:TeacherId;"`
}

func (classroomTeacher *ClassroomTeacher) TableName() string {
	return "classroom_teacher"
}

func (classroomTeacher *ClassroomTeacher) ToString() string {
	return fmt.Sprintf("ClassroomId: %d TeacherId: %d\n",
		classroomTeacher.ClassroomId, classroomTeacher.TeacherId)
}
