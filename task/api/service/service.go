package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mailchimp/common"
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

	create, err := common.PostCampaign(url, jsonData, apiKey)
	if err != nil {
		return nil, err
	}

	return create, nil
}

func GetCampaigns() ([]byte, error) {

	apiKey, serverPrefix, err := config.LoadEnv()

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns", serverPrefix)

	get, err := common.Get(url, apiKey)
	if err != nil {
		return nil, err
	}
	return get, nil

}

func GetCampaignsById(campaignid string) ([]byte, error) {

	apiKey, serverPrefix, err := config.LoadEnv()

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns/%s", serverPrefix, campaignid)

	get, err := common.GetById(url, apiKey)
	if err != nil {
		return nil, err
	}

	return get, nil
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

	update, err := common.UpdateById(url, jsonData, apiKey)
	if err != nil {
		return nil, err
	}

	return update, nil
}

func DeleteCampaignById(campaignID string) error {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns/%s", serverPrefix, campaignID)

	err = common.DeleteById(url, apiKey)
	if err != nil {
		return err
	}
	return nil
}

// func SendCampaign(campaignid string) error {
// 	apiKey, serverPrefix, err := config.LoadEnv()
// 	if err != nil {
// 		return err
// 	}

// 	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns/%s/actions/send", serverPrefix, campaignid)

// 	httpReq, err := http.NewRequest("POST", url, nil)
// 	if err != nil {
// 		return err
// 	}

// 	httpReq.Header.Set("Content-Type", "application/json")
// 	httpReq.SetBasicAuth("anystring", apiKey)

// 	resp, err := http.DefaultClient.Do(httpReq)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusNoContent {
// 		return fmt.Errorf("failed to send campaign %v", err)
// 	}
// 	return nil
// }

// func SetTemplateService(campaignID string, templateID int) error {
// 	apiKey, serverPrefix, err := config.LoadEnv()
// 	if err != nil {
// 		return err
// 	}

// 	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns/%s/content", serverPrefix, campaignID)

// 	body := map[string]interface{}{
// 		"template": map[string]interface{}{
// 			"id": templateID,
// 		},
// 	}

// 	jsonData, _ := json.Marshal(body)

// 	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	req.SetBasicAuth("anystring", apiKey)

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	respBody, _ := io.ReadAll(resp.Body)

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("failed to set template: status %d, body: %s", resp.StatusCode, string(respBody))
// 	}

// 	fmt.Println(" Template set successfully for campaign:", campaignID)
// 	return nil
// }

// Create a campaign using pre-made template
func CreateCampaign(req pkg.CampaignCreateRequest) (*pkg.CampaignResponse, error) {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns", serverPrefix)
	jsonData, _ := json.Marshal(req)

	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth("anystring", apiKey)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Mailchimp create campaign error: %s\n%s", resp.Status, string(body))
	}

	var campaignResp pkg.CampaignResponse
	_ = json.Unmarshal(body, &campaignResp)
	return &campaignResp, nil
}

// Send campaign by ID new
func SendCampaignService(campaignID string) error {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns/%s/actions/send", serverPrefix, campaignID)
	httpReq, _ := http.NewRequest("POST", url, nil)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth("anystring", apiKey)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Mailchimp send campaign error: %d\n%s", resp.StatusCode, string(body))
	}

	return nil
}

func SendCampaign(campaignID string) error {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/campaigns/%s/actions/send", serverPrefix, campaignID)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("anystring", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to send campaign: status %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

// Create Audience
func CreateAudienceService(req pkg.AudienceRequest) ([]byte, error) {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists", serverPrefix)
	data, _ := json.Marshal(req)

	create, err := common.Post(url, data, apiKey)
	if err != nil {
		return nil, err
	}
	return create, nil
}

// Get All Audiences
func GetAudiencesService() ([]byte, error) {
	apiKey, serverPrefix, _ := config.LoadEnv()
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists", serverPrefix)

	get, err := common.Get(url, apiKey)
	if err != nil {
		return nil, err
	}
	return get, nil
}

// Get Audience by ID
func GetAudienceByIdService(listID string) ([]byte, error) {
	apiKey, serverPrefix, _ := config.LoadEnv()
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s", serverPrefix, listID)

	get, err := common.GetById(url, apiKey)
	if err != nil {
		return nil, err
	}

	return get, nil
}

// Update Audience
func UpdateAudienceService(listID string, req pkg.AudienceRequest) ([]byte, error) {
	apiKey, serverPrefix, _ := config.LoadEnv()
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s", serverPrefix, listID)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	update, err := common.UpdateById(url, data, apiKey)
	if err != nil {
		return nil, err
	}

	return update, nil
}

// Delete Audience
func DeleteAudienceService(listID string) error {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s", serverPrefix, listID)

	err = common.DeleteById(url, apiKey)
	if err != nil {
		return err
	}
	return nil
}

// Create Member
func CreateMemberService(listID string, req pkg.MemberRequest) ([]byte, error) {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s/members", serverPrefix, listID)
	data, _ := json.Marshal(req)

	create, err := common.Post(url, data, apiKey)
	if err != nil {
		return nil, err
	}

	return create, nil
}

// Get All Members
func GetMembersService(listID string) ([]byte, error) {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s/members", serverPrefix, listID)

	get, err := common.Get(url, apiKey)
	if err != nil {
		return nil, err
	}
	return get, nil
}

// Get Member by Email (uses subscriber hash)
func GetMemberByEmailService(listID, email string) ([]byte, error) {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return nil, err
	}

	hash := common.ComputeSubscriberHash(email)
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s/members/%s", serverPrefix, listID, hash)

	get, err := common.GetById(url, apiKey)
	if err != nil {
		return nil, err
	}
	return get, nil
}

// Update Member (uses subscriber hash)
func UpdateMemberService(listID, email string, req pkg.MemberRequest) ([]byte, error) {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return nil, err
	}

	hash := common.ComputeSubscriberHash(email)
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s/members/%s", serverPrefix, listID, hash)

	data, _ := json.Marshal(req)
	update, err := common.UpdateById(url, data, apiKey)
	if err != nil {
		return nil, err
	}
	return update, nil
}

// Delete Member (uses subscriber hash)
func DeleteMemberService(listID, email string) error {
	apiKey, serverPrefix, err := config.LoadEnv()
	if err != nil {
		return err
	}

	hash := common.ComputeSubscriberHash(email)
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s/members/%s", serverPrefix, listID, hash)

	del := common.DeleteById(url, apiKey)
	return del
}
