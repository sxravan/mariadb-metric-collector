package config

import (
	"log"
	"os"
	"path/filepath"

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
	rootPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}
	envPath := filepath.Join(rootPath, ".env")
	err = godotenv.Load(envPath)

	if err != nil {
		log.Fatalf("Error loading .env file from %s: %v", envPath, err)
	}

	return &Config{
		DBUser:     os.Getenv("DBUser"),
		DBPassword: os.Getenv("DBPassword"),
		DBHost:     os.Getenv("DBHost"),
		DBName:     "information_schema",
		DBPort:     os.Getenv("DBPort"),
		DBSocket:   os.Getenv("DBSocket"),
	}, nil
}
