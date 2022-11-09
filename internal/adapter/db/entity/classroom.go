package entity

import (
	"fmt"
)

type Classroom struct {
	Id    int    `gorm:"primaryKey"`
	Grade string `gorm:"uniqueIndex:Grade"`
}

func (classroom *Classroom) TableName() string {
	return "classroom"
}

func (classroom *Classroom) ToString() string {
	return fmt.Sprintf("Classroom Id: %d Grade: %s\n", classroom.Id, classroom.Grade)
}
