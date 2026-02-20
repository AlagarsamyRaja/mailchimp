package handlers

import (
	"encoding/json"
	"mailchimp/api/service"
	"mailchimp/pkg"
	"net/http"
)

func CreateCampaignHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var campaignReq pkg.CampaignCreateRequest
	err := json.NewDecoder(r.Body).Decode(&campaignReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	response, err := service.CreateCampaignService(campaignReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetCampaign(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetCampaigns()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func GetCampaignById(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetCampaignsById()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
