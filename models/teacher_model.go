package models

type Teacher struct {
	ID              string  `json:"id"`
	UserID          string  `json:"userId"`
	Name            string  `json:"name"`
	Phone           string  `json:"phone"`
	Salary          float64 `json:"salary"`
	LeavesTaken     int     `json:"leavesTaken"`
	LeavesRemaining int     `json:"leavesRemaining"`
}
