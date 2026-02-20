package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mailchimp/config"
	"mailchimp/models"
	"net/http"
)

func CreateCampaignService(req models.CampaignCreateRequest) (*models.CampaignResponse, error) {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns", serverPrefix)

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth("anystring", apiKey) // Mailchimp requires anystring:apikey

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var campaignResp models.CampaignResponse
	json.NewDecoder(resp.Body).Decode(&campaignResp)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create campaign, status: %s", resp.Status)
	}

	return &campaignResp, nil
}
