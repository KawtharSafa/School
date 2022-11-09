package model

type CountFilter struct {
	Classroom    string `query:"classroom"`
	TeacherCount int
	StudentCount int
}

type ClassroomCounter struct {
	Classroom    string
	TeacherCount int
	StudentCount int
}
