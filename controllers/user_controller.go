package controllers

import (
	"go-library/models"
	"go-library/repository"
	"go-library/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo *repository.UserRepository
}

func NewUserController(repo *repository.UserRepository) *UserController {
	return &UserController{repo}
}

// GetUsers retrieves all users
func (uc *UserController) GetUsers(c echo.Context) error {
	page := 1

	if c.QueryParam("page") != "" {
		page = utils.QueryParamInt(c.QueryParam("page"), 1)
	}

	users, err := uc.repo.GetUsers(page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get users"})
	}
	return c.JSON(http.StatusOK, users)
}

// GetUserByID retrieves a user by their ID
func (uc *UserController) GetUserByID(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))

	user, err := uc.repo.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
func (uc *UserController) CreateUser(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "invalid user payload")
	}

	createdUser, err := uc.repo.CreateUser(user)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create user")
	}
	return c.JSON(http.StatusCreated, createdUser)
}

// UpdateUser updates a user
func (uc *UserController) UpdateUser(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "invalid user payload")
	}

	id := uuid.MustParse(c.Param("id"))

	updatedUser, err := uc.repo.UpdateUser(id, user)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "failed to update user")
	}
	return c.JSON(http.StatusOK, updatedUser)
}

// GetCheckoutsForUserByID retrieves all checkouts for a user
func (uc *UserController) GetCheckoutsForUserByID(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))

	checkouts, err := uc.repo.GetUsersCheckoutBooks(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get checkouts"})
	}
	return c.JSON(http.StatusOK, checkouts)
}

func (uc *UserController) CheckoutBook(c echo.Context) error {
	userId := uuid.MustParse(c.Param("id"))
	copyId := uuid.MustParse(c.Param("copy_id"))

	checkout, err := uc.repo.CheckoutBook(userId, copyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, checkout)
}

func (uc *UserController) ReturnBook(c echo.Context) error {
	userId := uuid.MustParse(c.Param("id"))
	copyId := uuid.MustParse(c.Param("copy_id"))

	err := uc.repo.ReturnBook(userId, copyId)
	if err !=	nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "book returned"})
}