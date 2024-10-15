package routes

import (
	"database/sql"
	"log"

	"go-library/controllers"
	"go-library/repository"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func SetupRoutes(e *echo.Echo, dbURL string) {
	// Initialize the database connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalln(err)
	}
	// defer db.Close()

	// Test the database connection
	if err := db.Ping(); err != nil {
		log.Fatalln("Could not connect to database:", err)
	}

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// Initialize repositories
	bookRepo := repository.NewBookRepository(db,)
	userRepo := repository.NewUserRepository(db)

	// Initialize controllers
	bookController := controllers.NewBookController(bookRepo)
	userController := controllers.NewUserController(userRepo)

	// Register routes
	BookRoutes(e, bookController)
	UserRoutes(e, userController)

	// list all routes
	e.GET("/routes", func(c echo.Context) error {
		routes := e.Routes()
		return c.JSON(200, routes)
	})

}
