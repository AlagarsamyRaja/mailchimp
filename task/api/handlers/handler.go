package handlers

import (
	"encoding/json"
	"mailchimp/api/service"
	"mailchimp/pkg"
	"net/http"

	"github.com/gorilla/mux"
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
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	campaignID := vars["id"]

	if campaignID == "" {
		http.Error(w, "Missing campaign ID", http.StatusBadRequest)
		return
	}

	data, err := service.GetCampaignsById(campaignID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UpdateCampaignHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	campaignID := vars["id"]

	if campaignID == "" {
		http.Error(w, "Missing campaign ID in URL", http.StatusBadRequest)
		return
	}

	var campaignReq pkg.CampaignCreateRequest
	err := json.NewDecoder(r.Body).Decode(&campaignReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data, err := service.UpdateCampaignService(campaignID, campaignReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func DeleteCampaignHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	campaignID := vars["id"]

	if campaignID == "" {
		http.Error(w, "Missing campaign ID in URL", http.StatusBadRequest)
		return
	}

	err := service.DeleteCampaignById(campaignID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := pkg.Response{Message: "Campaign deleted successfully"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func SendCampaign(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	campaignID := r.URL.Query().Get("id")
	if campaignID == "" {
		http.Error(w, "Campaign ID required", http.StatusBadRequest)
		return
	}

	err := service.SendCampaign(campaignID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Campaign sent successfully"))
}

// //audience
func CreateAudienceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req pkg.AudienceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data, err := service.CreateAudienceService(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// ✅ Get All Audiences
func GetAudiencesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := service.GetAudiencesService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// ✅ Get Audience by ID
func GetAudienceByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Missing audience ID", http.StatusBadRequest)
		return
	}

	data, err := service.GetAudienceByIdService(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

// ✅ Update Audience
func UpdateAudienceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Missing audience ID", http.StatusBadRequest)
		return
	}

	var req pkg.AudienceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data, err := service.UpdateAudienceService(id, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// ✅ Delete Audience
func DeleteAudienceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Missing audience ID", http.StatusBadRequest)
		return
	}

	err := service.DeleteAudienceService(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(pkg.Response{Message: "Audience deleted successfully"})
}
