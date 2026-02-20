package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mailchimp/config"
	"mailchimp/pkg"
	"net/http"

)

func CreateCampaignService(req pkg.CampaignCreateRequest) (*pkg.CampaignResponse, error) {
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
	httpReq.SetBasicAuth("anystring", apiKey)

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))

	var campaignResp pkg.CampaignResponse

	err=json.Unmarshal(body,&campaignResp)
	if err!=nil{
		return nil,err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create campaign, status: %s", resp.Status)
	}

	return &campaignResp, nil
}

func GetCampaigns() ([]byte, error) {

	apiKey, serverPrefix, err := config.LoadEnv()

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns", serverPrefix)

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("anystring", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func GetCampaignsById(campaignid string) ([]byte, error) {

	apiKey, serverPrefix, err := config.LoadEnv()

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns/%s", serverPrefix,campaignid)

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("anystring", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func DeleteCampaign(campaignid string) error {

	apiKey, serverPrefix, err := config.LoadEnv()

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns/%s", serverPrefix,campaignid)

	req, _ := http.NewRequest("DELETE", url, nil)

	req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func SendCampaign(campaignid string) error {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns/%s/actions/send", serverPrefix,campaignid)

	httpReq, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth("anystring", apiKey)

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent{
		return fmt.Errorf("failed to send campaign %v",err)
	}
	return nil
}