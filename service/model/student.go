package model

type StudentModel struct {
	FirstName string
	LastName  string
	Score     float32
	Classroom ClassroomModel
}
