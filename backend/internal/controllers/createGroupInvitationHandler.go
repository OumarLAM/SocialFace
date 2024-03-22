package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/models"
)

func InviteUserToGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	var inviteRequest struct {
		GroupID int `json:"group_id"`
		UserID  int `json:"user_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&inviteRequest)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
	}

	if !models.UserExists(inviteRequest.UserID) {
		http.Error(w, "User to invite does not exist", http.StatusBadRequest)
        return
	}

	isMember, err := models.IsGroupMember(userID, inviteRequest.GroupID)
	if err != nil {
        http.Error(w, "Failed to check if user is a member of the group", http.StatusInternalServerError)
        return
    }
	if !isMember {
		http.Error(w, "Only members of the group can invite others", http.StatusForbidden)
		return
	}

	err = models.SendGroupInvitation(inviteRequest.GroupID, userID, inviteRequest.UserID)
	if err != nil {
		http.Error(w, "Failed to send group invitation", http.StatusInternalServerError)
        return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Group invitation sent successfully"})
}

func AcceptInvitationToGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
        return
	}

	var acceptRequest struct {
       InvitationID int `json:"invitation_id"`
    }
	err := json.NewDecoder(r.Body).Decode(&acceptRequest)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
	}

	err = models.AcceptGroupInvitation(acceptRequest.InvitationID, userID)
	if err != nil {
		http.Error(w, "Failed to accept group invitation", http.StatusInternalServerError)
        return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Group invitation accepted successfully"})
}

func DeclineInvitationToGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
        return
	}

	var declineRequest struct {
		InvitationID int `json:"invitation_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&declineRequest)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
	}

	err = models.DeclineGroupInvitation(declineRequest.InvitationID, userID)
	if err != nil {
		http.Error(w, "Failed to decline group invitation", http.StatusInternalServerError)
        return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Group invitation declined successfully"})
}
