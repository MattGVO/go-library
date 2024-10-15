package routes

import (
	"go-library/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes (e *echo.Echo, controller *controllers.UserController) {
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUserByID)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)

	e.GET("/users/:id/checkouts", controller.GetCheckoutsForUserByID)
	e.POST("/users/:id/checkouts/copy/:copy_id", controller.CheckoutBook)
	e.PUT("/users/:id/checkouts/copy/:copy_id/return", controller.ReturnBook)
} 