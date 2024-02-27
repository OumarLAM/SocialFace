package middlewares

import (
    "net/http"

    "github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get session token from cookie
        sessionCookie, err := r.Cookie("session_token")
        if err != nil {
            http.Error(w, "Unhauthorized", http.StatusUnauthorized)
            return
        }
        sessionToken := sessionCookie.Value

        // Validate the session token against the database
        if !sqlite.IsSessionTokenValid(sessionToken) {
            http.Error(w, "Unhauthorized", http.StatusUnauthorized)
            return
        }

        // If the token is valid, call the next handler
        next.ServeHTTP(w, r)
    })
}
