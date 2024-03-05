package middlewares

import (
	"context"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Get session token from cookie
        sessionCookie, err := r.Cookie("session_token")
        if err != nil {
            http.Error(w, "Session token not found", http.StatusUnauthorized)
            return
        }
        sessionToken := sessionCookie.Value

        // Validate the session token against the database
        userID, valid := sqlite.IsSessionTokenValid(sessionToken)
        if !valid {
            http.Error(w, "Invalid session token", http.StatusUnauthorized)
            return
        }

        // Save userID in request context
        contex := context.WithValue(r.Context(), "userID", userID)

        // If the token is valid, call the next handler
        next.ServeHTTP(w, r.WithContext(contex))
    }
}
