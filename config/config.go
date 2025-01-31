package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
	DBPort     string
	DBSocket   string
}

func LoadConfig() (*Config, error) {
	// Load .env file if present
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}

	return &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     "information_schema",
		DBPort:     os.Getenv("DB_PORT"),
		DBSocket:   os.Getenv("DB_SOCKET"),
	}, nil
}
