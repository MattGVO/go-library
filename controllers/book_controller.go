package controllers

import (
	"go-library/models"
	"go-library/repository"
	"go-library/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	repo *repository.BookRepository
}

func NewBookController(repo *repository.BookRepository) *BookController {
	return &BookController{repo}
}

// GetBooks retrieves all books
func (bc *BookController) GetBooks(c echo.Context) error {
	page := 1

	if c.QueryParam("page") != "" {
		page = utils.QueryParamInt(c.QueryParam("page"), 1)
	}


	books, err := bc.repo.GetBooks(page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get books"})
	}
	return c.JSON(http.StatusOK, books)
}

// GetBookByID retrieves a book by its ID
func (bc *BookController) GetBookByID(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))

	book, err := bc.repo.GetBookByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "book not found"})
	}
	return c.JSON(http.StatusOK, book)
}

// CreateBook creates a new book
func (bc *BookController) CreateBook(c echo.Context) error {
	var book models.Book

	if err := c.Bind(&book); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "invalid book payload")
	}

	createdBook, err := bc.repo.CreateBook(book)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create book")
	}
	return c.JSON(http.StatusCreated, createdBook)
}

// UpdateBook updates a book
func (bc *BookController) UpdateBook(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	var book models.Book

	if err := c.Bind(&book); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "invalid book payload")
	}

	updatedBook, err := bc.repo.UpdateBook(id, book)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "failed to update book")
	}
	return c.JSON(http.StatusOK, updatedBook)
}

// DeleteBook deletes a book
func (bc *BookController) DeleteBook(c echo.Context) error {
	id := c.Param("id")

	err := bc.repo.DeleteBook(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "failed to delete book")
	}
	return c.NoContent(http.StatusNoContent)
}

func (bc *BookController) CreateCopyForBookByID(c echo.Context) error {
	bookID := uuid.MustParse(c.Param("id"))
	var copy models.Copy

	if err := c.Bind(&copy); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "invalid copy payload")
	}

	createdCopy, err := bc.repo.CreateCopyForBookByID(bookID, copy)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create copy")
	}
	return c.JSON(http.StatusCreated, createdCopy)
}

func (bc *BookController) UpdateCopyForBookByID(c echo.Context) error {
	bookID := uuid.MustParse(c.Param("id"))
	copyID := uuid.MustParse(c.Param("copy_id"))
	var copy models.Copy

	if err := c.Bind(&copy); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "invalid copy payload")
	}

	updatedCopy, err := bc.repo.UpdateCopyForBookByID(bookID, copyID, copy)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "failed to update copy")
	}
	return c.JSON(http.StatusOK, updatedCopy)
}