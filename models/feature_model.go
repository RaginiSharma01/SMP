package models

import "time"

type Classroom struct {
	ID             string `json:"id"`
	ClassName      string `json:"className"`
	Section        string `json:"section"`
	ClassTeacherID string `json:"classTeacherId"`
}
type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TeacherSubject struct {
	ID        string `json:"id"`
	TeacherID string `json:"teacherId"`
	SubjectID string `json:"subjectId"`
}
type Marks struct {
	ID         string `json:"id"`
	StudentID  string `json:"studentId"`
	SubjectID  string `json:"subjectId"`
	Term1Marks int    `json:"term1Marks"`
	Term2Marks int    `json:"term2Marks"`
}
type Timetable struct {
	ID          string `json:"id"`
	ClassroomID string `json:"classroomId"`
	SubjectID   string `json:"subjectId"`
	TeacherID   string `json:"teacherId"`
	DayOfWeek   string `json:"dayOfWeek"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}
type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	CreatedBy   string    `json:"createdBy"`
}
