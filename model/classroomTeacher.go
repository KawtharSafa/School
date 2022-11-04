package model

type ClassroomTeacherModel struct {
	Teacher      TeacherModel
	Classroom    ClassroomModel
	TeacherCount int
	StudentCount int
}
