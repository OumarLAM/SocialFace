package controllers

import (
    "encoding/json"
    "net/http"

	"github.com/OumarLAM/SocialFace/internal/models"
)

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body to get post ID
	var request struct {
		PostID int `json:"post_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
        http.Error(w, "User ID not found in context", http.StatusInternalServerError)
        return
    }

	err = models.LikePost(userID, request.PostID)
    if err != nil {
        http.Error(w, "Failed to like post", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post liked successfully"})
}

func UnlikePostHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		PostID int `json:"post_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
        http.Error(w, "User ID not found in context", http.StatusInternalServerError)
        return
    }

	err = models.UnlikePost(userID, request.PostID)
    if err != nil {
        http.Error(w, "Failed to unlike post", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Post unliked successfully"})
}