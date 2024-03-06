package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/models"
)

func FetchFollowingHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
    if !ok {
        http.Error(w, "User ID not found in context", http.StatusInternalServerError)
        return
    }

    // Fetch users followed by the user
    followers, err := models.FetchFollowing(userID)
    if err != nil {
        http.Error(w, "Failed to fetch followers", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(followers)
}