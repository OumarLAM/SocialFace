package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/models"
)

// GetProfileHandler retrieves user profile information based on the user ID.
func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from request context
	userId := getUserIDFromContext(r)

	// Fetch user profile information from the database
	user, err := models.GetUserByID(userId)
	if err != nil {
		http.Error(w, "Failed to fetch user profile", http.StatusInternalServerError)
		return
	}

	// Respond with user profile information
	w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

// UpdateProfilePrivacyHandler updates the profile privacy settings for the authenticated user
func UpdateProfilePrivacyHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from request context
    userId := getUserIDFromContext(r)

	// Parse request body
	var privacyUpdate struct {
		PublicProfile bool `json:"public_profile"`
	}
    err := json.NewDecoder(r.Body).Decode(&privacyUpdate)
    if err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    // Update user profile privacy settings in the database
    err = models.UpdateProfilePrivacy(userId, privacyUpdate.PublicProfile)
    if err != nil {
        http.Error(w, "Failed to update user profile privacy settings", http.StatusInternalServerError)
        return
    }

    // Respond with updated user profile privacy settings
    w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile privacy settings updated successfully"})
}