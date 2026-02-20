package config

import (
	"fmt"
	"mailchimp/pkg"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, string, string, error) {
	var datas pkg.MailchimpConfig
	err := godotenv.Load()
	if err != nil {
		return "", "", "", fmt.Errorf("error loading .env file: %v", err)
	}

	datas.ApiKey = os.Getenv("MAILCHIMP_API_KEY")
	datas.ServerPrefix = os.Getenv("MAILCHIMP_SERVER")

	// if datas.ApiKeyapiKey == "" || serverPrefix == "" {
	// 	return "", "", fmt.Errorf("missing required environment variables")
	// }

	return datas.ApiKey, datas.ServerPrefix, datas.CampaignId, nil
}
