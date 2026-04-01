package models

import "time"

type User struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	IsVerified bool      `json:"isVerified"`
	CreatedAt  time.Time `json:"createdAt"`
}
