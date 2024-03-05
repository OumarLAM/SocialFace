package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/models"
	"github.com/google/uuid"
)

// ProfileInfoHandler retrieves user profile information based on the user ID.
func ProfileInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Retrieve userID from the context
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
        return
	}
	
	// Fetch user profile information from the database based on the userID
	user, err := models.GetUserByID(userID)
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
	// Parse request body to get new privacy setting
	var privacySetting struct {
		PublicProfile bool `json:"public_profile"`
	}
	err := json.NewDecoder(r.Body).Decode(&privacySetting)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Extract userID from session token stored in the cookie
	userID, err := DecodeSessionToken(r)
	if err != nil {
		http.Error(w, "Failed to decode session token", http.StatusInternalServerError)
		return
	}

	// Update user profile privacy settings in the database
	err = models.UpdateProfilePrivacy(userID, privacySetting.PublicProfile)
	if err != nil {
		http.Error(w, "Failed to update user profile privacy settings", http.StatusInternalServerError)
		return
	}

	// Respond with updated user profile privacy settings
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile privacy settings updated successfully"})
}

// DecodeSessionToken decodes the session token to extract the userID.
func DecodeSessionToken(r *http.Request) (int, error) {
	// Extract the session token from cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, err
	}

	fmt.Println(cookie)

	// Parse session token as UUID
	sessionToken, err := uuid.Parse(cookie.Value)
	fmt.Println(sessionToken)
	if err != nil {
		return 0, err
	}

	// Extract the userID from the session token
	userID := int(sessionToken.ID())
	fmt.Println(userID)

	return userID, nil
}
