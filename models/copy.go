package models

import (
	"github.com/google/uuid"
)

type Copy struct {
	CopyID        uuid.UUID `json:"copy_id" db:"copy_id"`
	BookID        uuid.UUID `json:"book_id" db:"book_id"`
	AcquiredDate   string    `json:"acquired_date" db:"acquired_date"`
	Edition	   string    `json:"edition" db:"edition"`
	Condition string `json:"condition" db:"condition"`
}

type CopyWithCheckoutInfo struct {
	Copy
	CheckedOut bool `json:"checked_out"`
}

type CopyWithBookInfo struct {
	Copy
	Book Book `json:"book"`
}