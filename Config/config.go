package config

import (
	"fmt"
	"os"
)

func Config(key string) string {
	// load .env file
	_, err := os.ReadFile(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
