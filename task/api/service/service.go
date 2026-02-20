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
		return nil, err
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))

	var campaignResp pkg.CampaignResponse
	err = json.Unmarshal(body, &campaignResp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Mailchimp error: %s\nResponse: %s", resp.Status, string(body))
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

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns/%s", serverPrefix, campaignid)

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("anystring", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func UpdateCampaignService(campaignID string, req pkg.CampaignCreateRequest) ([]byte, error) {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns/%s", serverPrefix, campaignID)

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
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
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Mailchimp error: %s\nResponse: %s", resp.Status, string(body))
	}

	return body, nil
}

func DeleteCampaignById(campaignID string) error {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns/%s", serverPrefix, campaignID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth("anystring", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent { // 204 = Success (no content)
		return fmt.Errorf("failed to delete campaign, status: %s", resp.Status)
	}

	return nil
}

func SendCampaign(campaignid string) error {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns/%s/actions/send", serverPrefix, campaignid)

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

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to send campaign %v", err)
	}
	return nil
}

// ✅ Create Audience
func CreateAudienceService(req pkg.AudienceRequest) ([]byte, error) {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists", serverPrefix)
	data, _ := json.Marshal(req)

	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	httpReq.SetBasicAuth("anystring", apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Mailchimp error: %s\nResponse: %s", resp.Status, string(body))
	}
	return body, nil
}

// ✅ Get All Audiences
func GetAudiencesService() ([]byte, error) {
	apiKey, serverPrefix, _ := config.LoadEnv()
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists", serverPrefix)

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("anystring", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// ✅ Get Audience by ID
func GetAudienceByIdService(listID string) ([]byte, error) {
	apiKey, serverPrefix, _ := config.LoadEnv()
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s", serverPrefix, listID)

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("anystring", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// ✅ Update Audience
func UpdateAudienceService(listID string, req pkg.AudienceRequest) ([]byte, error) {
	apiKey, serverPrefix, _ := config.LoadEnv()
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s", serverPrefix, listID)

	data, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(data))
	httpReq.SetBasicAuth("anystring", apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Mailchimp error: %s\nResponse: %s", resp.Status, string(body))
	}
	return body, nil
}

// ✅ Delete Audience
func DeleteAudienceService(listID string) error {
	apiKey, serverPrefix, _ := config.LoadEnv()
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s", serverPrefix, listID)

	req, _ := http.NewRequest("DELETE", url, nil)
	req.SetBasicAuth("anystring", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete audience, status: %s, response: %s", resp.Status, string(body))
	}
	return nil
}
