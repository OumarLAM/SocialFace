package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/models"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if event.Title == "" || event.Description == "" || event.DateTime == "" || event.GroupID == 0 {
		http.Error(w, "Incomplete event data", http.StatusBadRequest)
		return
	}

	err := models.CreateEvent(&event)
	if err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "event created successfully"})
}

func RespondToEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var eventResponse models.EventResponse
	if err := json.NewDecoder(r.Body).Decode(&eventResponse); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if eventResponse.EventID == 0 || eventResponse.UserID == 0 || eventResponse.Response == "" {
		http.Error(w, "Incomplete event response data", http.StatusBadRequest)
		return
	}

	err := models.RespondToEvent(&eventResponse)
	if err != nil {
		http.Error(w, "Failed to respond to event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "event response recorded successfully"})
}
