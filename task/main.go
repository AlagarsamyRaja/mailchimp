package main

import (
	"fmt"
	"mailchimp/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/createcampaign", handlers.CreateCampaignHandler).Methods("POST")
	r.HandleFunc("/getcampaign", handlers.GetCampaign).Methods("GET")
	r.HandleFunc("/getcampaignbyid/{id}", handlers.GetCampaignById).Methods("GET")
	r.HandleFunc("/updatecampaign/{id}", handlers.UpdateCampaignHandler).Methods("PATCH")
	r.HandleFunc("/deletecampaignbyid/{id}", handlers.DeleteCampaignHandler).Methods("DELETE")

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
