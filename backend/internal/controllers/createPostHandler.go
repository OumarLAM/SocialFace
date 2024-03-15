package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/models"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if post.Content == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	if (post.Privacy != models.Public && post.Privacy != models.Private && post.Privacy != models.AlmostPrivate) {
		http.Error(w, "Invalid privacy level", http.StatusBadRequest)
        return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	err = models.CreatePost(userID, post)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Post created successfully"})
}
