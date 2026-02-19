package config

import (
	"fmt"
	"mailchimp/pkg"
	"os"

	"github.com/joho/godotenv"
)

func MailchimpConfiguration() (pkg.MailchimpConfig, error) {
	var data pkg.MailchimpConfig

	err := godotenv.Load()
	if err != nil {
		return data, fmt.Errorf("Error loading .env file: %v", err)
	}

	data.ApiKey = os.Getenv("MAILCHIMP_API_KEY")
	data.ServerPrefix = os.Getenv("MAILCHIMP_SERVER_PREFIX")
	data.ListID = os.Getenv("MAILCHIMP_LIST_ID")

	if data.ApiKey == "" || data.ServerPrefix == "" {
		return data, fmt.Errorf("Missing Mailchimp configuration values")
	}

	return data, nil
}
