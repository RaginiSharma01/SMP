package models

import "time"

type Student struct {
	ID           string    `json:"id"` 
	RollNumber   int       `json:"rollNumber"`
	ClassroomID  string    `json:"classroomId"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Phone        string    `json:"phone"`
	DOB          time.Time `json:"dob"`
	Age          int       `json:"age"`
	Address      string    `json:"address"`
	FatherName   string    `json:"fatherName"`
	MotherName   string    `json:"motherName"`
	GuardianName string    `json:"guardianName"`
	Occupation   string    `json:"occupation"`
	Height       float64   `json:"height"`
	Weight       float64   `json:"weight"`
	PhotoURL     string    `json:"photoUrl"`
}
