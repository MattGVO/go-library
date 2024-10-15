package routes

import (
	"go-library/controllers"

	"github.com/labstack/echo/v4"
)

func BookRoutes(e *echo.Echo, controller *controllers.BookController) {
	e.GET("/books", controller.GetBooks)
	e.GET("/books/:id", controller.GetBookByID)
	e.POST("/books", controller.CreateBook)
	e.PUT("/books/:id", controller.UpdateBook)
	e.DELETE("/books/:id", controller.DeleteBook)

	e.POST("/books/:id/copy", controller.CreateCopyForBookByID)
	e.PUT("/books/:id/copy/:copy_id", controller.UpdateCopyForBookByID)
}
