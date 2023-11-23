package methods

import (
	"crypto/subtle"
	"net/http"
	"os"
)

// basicAuth - adding basic authentication to route
func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			username := []byte(username)
			password := []byte(password)
			expectedUsername := []byte(os.Getenv("APP_USERNAME"))
			expectedPassword := []byte(os.Getenv("APP_PASSWORD"))

			usernameMatch := (subtle.ConstantTimeCompare(username[:], expectedUsername[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(password[:], expectedPassword[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
