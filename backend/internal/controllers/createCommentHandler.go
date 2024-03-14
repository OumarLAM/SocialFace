package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/models"
	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

	if comment.Content == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
        return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
        return
	}

	if !sqlite.PostExists(comment.PostID) {
		http.Error(w, "Post not found", http.StatusNotFound)
        return
	}

	err = sqlite.CreateComment(userID, comment)
	if err != nil {
        http.Error(w, "Failed to create comment", http.StatusInternalServerError)
        return
    }

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Comment created successfully"})
}