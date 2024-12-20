package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetDotenvVariable(variable string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Could not load .env file")
	}

	return os.Getenv(variable)
}
