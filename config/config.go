package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds application configuration
type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	ServerPort	   string
}

// LoadConfig loads configuration from .env file
func LoadConfig() (Config, error) {
	var config Config

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config.PostgresUser = os.Getenv("POSTGRES_USER")
	config.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	config.PostgresDB = os.Getenv("POSTGRES_DB")
	config.ServerPort = os.Getenv("SERVER_PORT")

	return config, nil
}
