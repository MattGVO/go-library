package models

import (
	"github.com/google/uuid"
)

type Book struct {
	BookID        uuid.UUID `json:"book_id" db:"book_id"`
	Title         string    `json:"title" db:"title"`
	Author        string    `json:"author" db:"author"`
	Genre         string    `json:"genre" db:"genre"`
	PublishedYear int       `json:"published_year" db:"published_year"`
}

type BookWithCopiesAndCheckoutInfo struct {
	Book
	Copies []CopyWithCheckoutInfo `json:"copies"`
}
