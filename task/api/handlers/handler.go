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

	campaignid:=r.URL.Query().Get("id")

	data, err := service.GetCampaignsById(campaignid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func DeleteCampaignHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    campaignID := r.URL.Query().Get("id")
    if campaignID == "" {
        http.Error(w, "Campaign ID is required", http.StatusBadRequest)
        return
    }

    err := service.DeleteCampaign(campaignID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte("Campaign deleted successfully"))
}