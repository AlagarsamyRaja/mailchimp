package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mailchimp/pkg"
	"net/http"
)

func PostCampaign(url string, jsonData []byte,apikey string) (*pkg.CampaignResponse, error){
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth("anystring", apikey)

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


func Post(url string,data []byte,apikey string)([]byte, error){
	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	httpReq.SetBasicAuth("anystring", apikey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Error: %s\nResponse: %s", resp.Status, string(body))
	}
	return body, nil
}

func Get(url string,apikey string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("anystring", apikey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func GetById(url string,apikey string)([]byte, error){
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("anystring", apikey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func UpdateById(url string,jsonData []byte,apikey string)([]byte, error) {

	httpReq, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth("anystring", apikey)

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


func DeleteById(url string,apikey string) error{
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth("anystring", apikey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent { 
		return fmt.Errorf("failed to delete campaign, status: %s", resp.Status)
	}

	return nil
}