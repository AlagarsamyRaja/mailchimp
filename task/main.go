package main

import (
	"fmt"
	"mailchimp/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/createcampaign", handlers.CreateCampaigns).Methods("POST")
	r.HandleFunc("/getcampaign", handlers.GetCampaign).Methods("GET")
	r.HandleFunc("/getcampaignbyid/{id}", handlers.GetCampaignById).Methods("GET")
	r.HandleFunc("/updatecampaign/{id}", handlers.UpdateCampaignHandler).Methods("PATCH")
	r.HandleFunc("/deletecampaignbyid/{id}", handlers.DeleteCampaignHandler).Methods("DELETE")
	//r.HandleFunc("/setcampaign/{id}", handlers.SetTemplateHandler).Methods("PUT")
	r.HandleFunc("/sendcampaign/{id}", handlers.SendCampaignHandler).Methods("POST")

	r.HandleFunc("/createaudience", handlers.CreateAudienceHandler).Methods("POST")
	r.HandleFunc("/getaudience", handlers.GetAudiencesHandler).Methods("GET")
	r.HandleFunc("/getaudiencebyid/{id}", handlers.GetAudienceByIdHandler).Methods("GET")
	r.HandleFunc("/updateaudience/{id}", handlers.UpdateAudienceHandler).Methods("PATCH")
	r.HandleFunc("/deleteaudiencebyid/{id}", handlers.DeleteAudienceHandler).Methods("DELETE")

	r.HandleFunc("/createmember/{id}", handlers.CreateMemberHandler).Methods("POST")
	r.HandleFunc("/getmember/{id}", handlers.GetMembersHandler).Methods("GET")
	r.HandleFunc("/list/{id}/members/{email}", handlers.GetMemberByEmailHandler).Methods("GET")
	r.HandleFunc("/list/{id}/members/{email}", handlers.UpdateMemberHandler).Methods("PATCH")
	r.HandleFunc("/list/{id}/members/{email}", handlers.DeleteMemberHandler).Methods("DELETE")

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
