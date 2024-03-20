package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/models"
)

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
	}

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

	postExists, err := models.PostExists(request.PostID)
	if err != nil {
        http.Error(w, "Failed to check if post exists", http.StatusInternalServerError)
        return
    }

	if !postExists {
		http.Error(w, "Post does not exist", http.StatusNotFound)
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

func DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
	}

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

	postExists, err := models.PostExists(request.PostID)
	if err != nil {
        http.Error(w, "Failed to check if post exists", http.StatusInternalServerError)
        return
    }

	if !postExists {
		http.Error(w, "Post does not exist", http.StatusNotFound)
        return
	}

	err = models.DislikePost(userID, request.PostID)
	if err != nil {
		http.Error(w, "Failed to dislike post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post disliked successfully"})
}
