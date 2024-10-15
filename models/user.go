package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID `json:"user_id" db:"user_id"`
	FullName string    `json:"full_name" db:"full_name"`
	Email    string    `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	RegisteredDate string `json:"registered_date" db:"registered_date"`
}