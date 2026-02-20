package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", "", fmt.Errorf("error loading .env file: %v", err)
	}

	apiKey := os.Getenv("MAILCHIMP_API_KEY")
	serverPrefix := os.Getenv("MAILCHIMP_SERVER_PREFIX")

	if apiKey == "" || serverPrefix == "" {
		return "", "", fmt.Errorf("missing required environment variables")
	}

	return apiKey, serverPrefix, nil
}
