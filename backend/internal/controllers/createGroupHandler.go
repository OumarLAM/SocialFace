package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/models"
)

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	var group models.UserGroup
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if group.Title == "" || group.Description == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	group.CreatorID = userID
	err := models.CreateGroup(group)
	if err != nil {
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Group created successfully"})
}

func LeaveGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
        return
	}

	var leaveRequest struct {
		GroupID int `json:"group_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&leaveRequest)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
	}

	err = models.LeaveGroup(userID, leaveRequest.GroupID)
	if err != nil {
		http.Error(w, "Failed to leave group", http.StatusInternalServerError)
        return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Group left successfully"})
}
