package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Checkout struct {
	CheckoutID    uuid.UUID `json:"checkout_id" db:"checkout_id"`
	CopyID        uuid.UUID `json:"copy_id" db:"copy_id"`
	UserID        uuid.UUID `json:"user_id" db:"user_id"`
	CheckoutDate  string    `json:"checkout_date" db:"checkout_date"`
	ReturnDate    sql.NullString    `json:"return_date" db:"return_date"`
	DueDate       string    `json:"due_date" db:"due_date"`
}

type CheckoutWithCopyAndBookInfo struct {
	Checkout
	CopyWithBookInfo CopyWithBookInfo `json:"copy"`
}