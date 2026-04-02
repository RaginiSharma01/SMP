package models

import "time"

type Classroom struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	TeacherID    string    `json:"teacherId"`
	Subjects     string    `json:"subjects"`
	AcademicYear string    `json:"academicYear"`
	CreatedAt    time.Time `json:"createdAt"`
}