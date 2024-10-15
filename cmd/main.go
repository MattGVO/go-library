package main

import (
	"go-library/config"
	"go-library/routes"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load config:", err)
	}

	// Create a new Echo instance
	e := echo.New()

	dbUrl := "postgres://" + cfg.PostgresUser + ":" + cfg.PostgresPassword + "@localhost/" + cfg.PostgresDB + "?sslmode=disable"

	// Setup routes
	routes.SetupRoutes(e, dbUrl)

	// Start the server
	e.Logger.Fatal(e.Start(cfg.ServerPort))
}
