package common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

func Post(url string, data []byte, apikey string) ([]byte, error) {
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

func Get(url string, apikey string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("anystring", apikey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func GetById(url string, apikey string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("anystring", apikey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func UpdateById(url string, jsonData []byte, apikey string) ([]byte, error) {

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

func DeleteById(url string, apikey string) error {
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

func ComputeSubscriberHash(email string) string {
	hash := md5.Sum([]byte(email))
	return hex.EncodeToString(hash[:])
}
