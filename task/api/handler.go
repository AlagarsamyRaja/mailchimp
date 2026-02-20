package api

import (
	"encoding/json"
	"fmt"
	"mailchimp/models"
	"mailchimp/service"
	"net/http"
)

func CreateCampaignHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var campaignReq models.CampaignCreateRequest
	err := json.NewDecoder(r.Body).Decode(&campaignReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	response, err := service.CreateCampaignService(campaignReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
