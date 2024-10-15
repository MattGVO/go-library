package repository

import (
	"database/sql"
	"fmt"

	"go-library/models"
	"go-library/utils"

	"github.com/google/uuid"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db}
}
// GetBooks retrieves all books from the database paginated
func (r *BookRepository) GetBooks(page int) ([]models.BookWithCopiesAndCheckoutInfo, error) {
	offset := (page - 1) * 10

	if page == 1 {
		offset = 0
	}

	rows, err := r.DB.Query("SELECT id, title, author, genre, published_year FROM books b LIMIT 10 OFFSET $1", offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice of books to be populated
	var books []models.Book

	
	// Use the utility function to scan rows into the books slice
	if err := utils.ScanRows(rows, &books); err != nil {
		return nil, err
	}
	// same length as books
	var bookCopies = make([]models.BookWithCopiesAndCheckoutInfo, len(books))

	for i, book := range books {
		rows, err := r.DB.Query("SELECT id, book_id, acquired_date, edition, condition FROM copies WHERE book_id = $1", book.BookID);

		if err != nil {
			return nil, err
		}

		defer rows.Close()

		copies := []models.Copy{}

		if err := utils.ScanRows(rows, &copies); err != nil {
			return nil, err
		}

		bookCopies[i].Book = book
		copiesWithCheckoutInfo := []models.CopyWithCheckoutInfo{}

		for _, copy := range copies {
			rows, err := r.DB.Query("SELECT c.id, c.copy_id, c.user_id, c.checkout_date, c.return_date, c.due_date FROM checkouts c WHERE c.copy_id = $1", copy.CopyID)
			if err != nil {
				return nil, err
			}

			defer rows.Close()

			var checkouts []models.Checkout

			if err := utils.ScanRows(rows, &checkouts); err != nil {
				return nil, err
			}

			copiesWithCheckoutInfo = append(copiesWithCheckoutInfo, models.CopyWithCheckoutInfo{
				Copy: copy,
				CheckedOut: len(checkouts) > 0,
			})
		}
		bookCopies[i].Copies = append(bookCopies[i].Copies, copiesWithCheckoutInfo...)
	
	}

	return bookCopies, nil
}

func (r *BookRepository) GetBookByID(id uuid.UUID) (models.Book, error) {
	var book models.Book

	row := r.DB.QueryRow("SELECT id, title, author, genre, published_year FROM books WHERE id = $1", id)
	if err := row.Scan(&book.BookID, &book.Title, &book.Author, &book.Genre, &book.PublishedYear); err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (r *BookRepository) CreateBook(book models.Book) (models.Book, error) {
	err := r.DB.QueryRow("INSERT INTO books (title, author, genre, published_year) VALUES ($1, $2, $3, $4) RETURNING id", book.Title, book.Author, book.Genre, book.PublishedYear).Scan(&book.BookID)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (r *BookRepository) UpdateBook(id uuid.UUID, book models.Book) (models.Book, error) {
    result, err := r.DB.Exec("UPDATE books SET title = $1, author = $2, genre = $3, published_year = $4 WHERE id = $5", 
        book.Title, book.Author, book.Genre, book.PublishedYear, id)
    if err != nil {
        return models.Book{}, err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return models.Book{}, err
    }

    if rowsAffected == 0 {
        return models.Book{}, fmt.Errorf("no rows were updated")
    }

    book.BookID = id
    return book, nil
}

func (r *BookRepository) DeleteBook(id string) error {
	_, errBook := r.DB.Exec("DELETE FROM books WHERE id = $1", id)

	if errBook != nil {
		return errBook
	}

	_, errCopies := r.DB.Exec("DELETE FROM copies WHERE book_id = $1", id)

	if errCopies != nil {
		return errCopies
	}
	
	return nil
}

func (r *BookRepository) CreateCopyForBookByID(bookID uuid.UUID, copy models.Copy) (models.Copy, error) {
	book, err := r.GetBookByID(bookID)
	if err != nil {
		return models.Copy{}, err
	}

	copy.BookID = book.BookID
	copy.CopyID = uuid.New()

	_, err = r.DB.Exec("INSERT INTO copies (id, book_id, acquired_date, edition, condition) VALUES ($1, $2, $3, $4, $5)",
		copy.CopyID, copy.BookID, copy.AcquiredDate, copy.Edition, copy.Condition)
	if err != nil {
		return models.Copy{}, err
	}


	return copy, nil
}

func (r *BookRepository) UpdateCopyForBookByID(bookID uuid.UUID, copyID uuid.UUID, copy models.Copy) (models.Copy, error) {
	book, err := r.GetBookByID(bookID)
	if err != nil {
		return models.Copy{}, err
	}

	copy.BookID = book.BookID
	copy.CopyID = copyID

	_, err = r.DB.Exec("UPDATE copies edition = $3, condition = $4 WHERE id = $5",
		copy.Edition, copy.Condition, copyID)
	if err != nil {
		return models.Copy{}, err
	}

	return copy, nil
}


