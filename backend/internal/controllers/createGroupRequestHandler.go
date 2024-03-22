package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/models"
)

func RequestToJoinGroupHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		UserID  int `json:"user_id"`
		GroupID int `json:"group_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = models.AddGroupRequest(request.UserID, request.GroupID)
	if err != nil {
		http.Error(w, "Failed to add group request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Group request sent successfully"})
}

func AcceptGroupRequestHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		UserID    int `json:"user_id"`
		GroupID   int `json:"group_id"`
		RequestID int `json:"request_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = models.AcceptGroupRequest(request.UserID, request.GroupID, request.RequestID)
	if err != nil {
		http.Error(w, "Failed to accept group request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Group request accepted successfully"})
}

func DeclineGroupRequestHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		UserID    int `json:"user_id"`
		GroupID   int `json:"group_id"`
		RequestID int `json:"request_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = models.DeclineGroupRequest(request.UserID, request.GroupID, request.RequestID)
	if err != nil {
		http.Error(w, "Failed to decline group request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Group request declined successfully"})
}
