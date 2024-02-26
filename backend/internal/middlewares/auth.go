package middlewares

import "net/http"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if user is authenticated
		// If not, redirect to login page or return unauthorized repspo
		// Otherwise, call next handler
		// session, err := store.Get(r, "session")
        // if err!= nil {
        //     http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
        //     return
        // }

        // if session.Values["authenticated"] == nil {
        //     http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
        //     return
        // }

        next(w, r)
	}
}