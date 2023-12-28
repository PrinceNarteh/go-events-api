package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
