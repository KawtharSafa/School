package entity

import (
	"fmt"
)

type Teacher struct {
	Id          int `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Description string
}

func (teacher *Teacher) TableName() string {
	return "teacher"
}

func (teacher *Teacher) ToString() string {
	return fmt.Sprintf("teacher Id: %d First Name: %s last Name: %s Description: %s\n",
		teacher.Id, teacher.FirstName, teacher.LastName, teacher.Description)
}
