package methods

import (
	"context"
	"crypto/subtle"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"go.uber.org/zap"
)

// Keycloak settings
type Keycloak struct {
	Logger       *zap.SugaredLogger
	Gocloak      *gocloak.GoCloak
	ClientId     string
	ClientSecret string
	Realm        string
	token        string
	Permissions  *[]gocloak.RequestingPartyPermission
}

func (auth *Keycloak) extractBearerToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}

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

// Validate access token. No detailed permissions
func (auth *Keycloak) VerifyToken(next http.HandlerFunc) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {

		// try to extract Authorization parameter from the HTTP header
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// extract Bearer token
		auth.token = auth.extractBearerToken(token)
		if token == "" {
			http.Error(w, "Bearer Token missing", http.StatusUnauthorized)
			return
		}

		fmt.Printf("token: %v\n", token)

		//// call Keycloak API to verify the access token
		result, err := auth.Gocloak.RetrospectToken(context.Background(), auth.token, auth.ClientId, auth.ClientSecret, auth.Realm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		// check that token is correct
		jwt, _, err := auth.Gocloak.DecodeAccessToken(context.Background(), auth.token, auth.Realm)
		if err != nil {
			fmt.Printf("token: %v\n", jwt)
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		// check if the token isn't expired and valid
		if !*result.Active {
			http.Error(w, "Invalid or expired Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}

// Validate access token. With detail permissions
func (auth *Keycloak) VerifyTokenPermissions(permissions string, next http.HandlerFunc) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {

		// try to extract Authorization parameter from the HTTP header
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// extract Bearer token
		auth.token = auth.extractBearerToken(token)
		if token == "" {
			http.Error(w, "Bearer Token missing", http.StatusUnauthorized)
			return
		}

		//// call Keycloak API to verify the access token
		result, err := auth.Gocloak.RetrospectToken(context.Background(), auth.token, auth.ClientId, auth.ClientSecret, auth.Realm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		// check that token is correct
		jwt, _, err := auth.Gocloak.DecodeAccessToken(context.Background(), auth.token, auth.Realm)
		if err != nil {
			fmt.Printf("token: %v\n", jwt)
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		audience := auth.ClientId
		options := gocloak.RequestingPartyTokenOptions{
			Audience: &audience,
		}

		// call Keycloak API for a access token with permissions
		auth.Permissions, err = auth.Gocloak.GetRequestingPartyPermissions(context.Background(), auth.token, auth.Realm, options)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		// checking for authorization, default no access
		permMap := strings.Split(permissions, ":")
		resource := permMap[0]
		scope := permMap[1]
		access := false

		for _, claim := range *auth.Permissions {
			if *claim.ResourceName == resource {
				for _, claimScope := range *claim.Scopes {
					if claimScope == scope {
						access = true
					}
				}
			}
		}

		// check if the token isn't expired and valid
		if !*result.Active {
			http.Error(w, "Invalid or expired Token", http.StatusUnauthorized)
			return
		}

		// all:get gives permissions to endpoints. Needs to be avoided
		if permissions == "all:get" {
			access = true
		}

		if !access {
			http.Error(w, "No access", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}
