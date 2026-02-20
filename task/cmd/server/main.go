package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"
// )

// //Create audience

// // Define the structure for the data expected from Postman
// type MailchimpMember struct {
// 	EmailAddress string `json:"email_address"`
// 	Status       string `json:"status"`
// 	MergeFields  struct {
// 		FNAME string `json:"FNAME"`
// 		LNAME string `json:"LNAME"`
// 	} `json:"merge_fields"`
// }

// // Handler function to process the request from Postman
// func addMemberHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// 1. Decode the JSON data received from Postman
// 	var memberData MailchimpMember
// 	err := json.NewDecoder(r.Body).Decode(&memberData)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// 2. Prepare the request to the Mailchimp API
// 	api_key := os.Getenv("MAILCHIMP_API_KEY")
// 	list_id := os.Getenv("MAILCHIMP_LIST_ID")
// 	server_prefix := os.Getenv("MAILCHIMP_SERVER_PREFIX")

// 	if api_key == "" || list_id == "" || server_prefix == "" {
// 		log.Fatal("MAILCHIMP environment variables not set")
// 	}

// 	url := fmt.Sprintf("https://%://s.api.mailchimp.com", server_prefix, list_id)

// 	jsonData, err := json.Marshal(memberData)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Create an HTTP client and request
// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// 3. Set the necessary headers, including Basic Authentication
// 	// The username for Mailchimp API is anything (usually your username or a blank string),
// 	// and the password is the API key.
// 	req.SetBasicAuth("anyuser", api_key)
// 	req.Header.Add("Content-Type", "application/json")

// 	// 4. Send the request to Mailchimp
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// 5. Handle the Mailchimp API response
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	if resp.StatusCode >= 400 {
// 		http.Error(w, fmt.Sprintf("Mailchimp API error: %s", string(body)), resp.StatusCode)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(fmt.Sprintf("Successfully added member to Mailchimp audience: %s", string(body))))
// }

// func main() {
// 	// Set your Mailchimp environment variables (API Key, List ID, Server Prefix) before running
// 	// Example: os.Setenv("MAILCHIMP_API_KEY", "your-api-key-usX")
// 	// Example: os.Setenv("MAILCHIMP_LIST_ID", "your-list-id")
// 	// Example: os.Setenv("MAILCHIMP_SERVER_PREFIX", "usX")

// 	http.HandleFunc("/add-member", addMemberHandler)
// 	fmt.Println("Server listening on port 8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

import (
	"fmt"
	"log"
	"mailchimp/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/create-campaign", handler.CreateCampaignHandler)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
