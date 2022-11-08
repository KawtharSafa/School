package entity

import (
	"fmt"
	"school.com/packages/service/entity"
)

type Student struct {
	Id          int `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Score       float32
	ClassroomId int
	Classroom   entity.Classroom `gorm:"foreignKey:ClassroomId"`
}

func (student *Student) TableName() string {
	return "student"
}

func (student *Student) ToString() string {
	return fmt.Sprintf("Student Id: %d First Name: %s last Name: %s Classroom Id: %d\n",
		student.Id, student.FirstName, student.LastName, student.ClassroomId)
}
