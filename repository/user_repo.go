package repository

import (
	"database/sql"
	"errors"

	"go-library/models"
	"go-library/utils"

	"github.com/google/uuid"
)

// UserRepository is a struct to represent the user repository
type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

// GetUsers retrieves all users from the database paginated
func (r *UserRepository) GetUsers(page int) ([]models.User, error) {
	offset := (page - 1) * 10

	if page == 1 {
		offset = 0
	}

	rows, err := r.DB.Query("SELECT id, full_name, email, phone_number, registered_date FROM users LIMIT 10 OFFSET $1", offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	if err := utils.ScanRows(rows, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID retrieves a user by their ID
func (r *UserRepository) GetUserByID(id uuid.UUID) (models.User, error) {
	var user models.User

	row := r.DB.QueryRow("SELECT id, full_name, email, phone_number, registered_date FROM users WHERE id = $1", id)
	if err := row.Scan(&user.UserID, &user.FullName, &user.Email, &user.PhoneNumber, &user.RegisteredDate); err != nil {
		return models.User{}, err
	}

	return user, nil
}

// CreateUser creates a new user
func (r *UserRepository) CreateUser(user models.User) (models.User, error) {
	user.UserID = uuid.New()

	_, err := r.DB.Exec("INSERT INTO users (id, full_name, email, phone_number, registered_date) VALUES ($1, $2, $3, $4, $5)",
		user.UserID, user.FullName, user.Email, user.PhoneNumber, user.RegisteredDate)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// UpdateUser updates a user
func (r *UserRepository) UpdateUser(id uuid.UUID, user models.User) (models.User, error) {
	_, err := r.DB.Exec("UPDATE users SET full_name = $1, email = $2, phone_number = $3 WHERE id = $4",
		user.FullName, user.Email, user.PhoneNumber, user.UserID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUsersCheckoutBooks(id uuid.UUID) ([]models.CheckoutWithCopyAndBookInfo, error) {
	// get checkouts for user
	rows, err := r.DB.Query("SELECT c.id, c.copy_id, c.user_id, c.checkout_date, c.return_date, c.due_date FROM checkouts c WHERE c.user_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var checkouts []models.Checkout

	if err := utils.ScanRows(rows, &checkouts); err != nil {
		return nil, err
	}

	checkoutsWithCopyAndBookInfo := make([]models.CheckoutWithCopyAndBookInfo, len(checkouts))

	for i, checkout := range checkouts {
		result := r.DB.QueryRow("SELECT id, book_id, acquired_date, edition, condition FROM copies WHERE id = $1", checkout.CopyID)
		if err != nil {
			return nil, err
		}

		var copy models.Copy
		if err := utils.ScanRow(result, &copy); err != nil {
			return nil, err
		}

		book := r.DB.QueryRow("SELECT id, title, author, genre, published_year FROM books WHERE id = $1", copy.BookID)
		if err != nil {
			return nil, err
		}

		var bookInfo models.Book
		if err := utils.ScanRow(book, &bookInfo); err != nil {
			return nil, err
		}

		checkoutsWithCopyAndBookInfo[i] = models.CheckoutWithCopyAndBookInfo{
			Checkout: checkout,
			CopyWithBookInfo: models.CopyWithBookInfo{
				Copy: copy,
				Book: bookInfo,
			},
		}

	}

	return checkoutsWithCopyAndBookInfo, nil
}

func (r *UserRepository) CheckoutBook(userID, copyID uuid.UUID) (models.Checkout, error) {
	checkoutID := uuid.New()

	// check for existing checkout
	row := r.DB.QueryRow("SELECT id, copy_id, user_id, checkout_date, return_date, due_date FROM checkouts WHERE copy_id = $1 AND return_date IS NULL", copyID)
	var checkout models.Checkout
	if err := row.Scan(&checkout.CheckoutID, &checkout.CopyID, &checkout.UserID, &checkout.CheckoutDate, &checkout.ReturnDate, &checkout.DueDate); err == nil {
		return models.Checkout{}, errors.New("copy is already checked out")
	}

	_, err := r.DB.Exec("INSERT INTO checkouts (id, copy_id, user_id, checkout_date, due_date) VALUES ($1, $2, $3, $4, $5)",
		checkoutID, copyID, userID, utils.GetCurrentDate(), utils.GetDueDate())
	if err != nil {
		return models.Checkout{}, err
	}

	return models.Checkout{
		CheckoutID:   checkoutID,
		CopyID:       copyID,
		UserID:       userID,
		CheckoutDate: utils.GetCurrentDate(),
		DueDate:      utils.GetDueDate(),
	}, nil
}

func (r *UserRepository) ReturnBook(userID, copyID uuid.UUID) error {
	_, err := r.DB.Exec("UPDATE checkouts SET return_date = $1 WHERE user_id = $2 AND copy_id = $3 AND return_date IS NULL",
		utils.GetCurrentDate(), userID, copyID)
	if err != nil {
		return err
	}

	return nil
}