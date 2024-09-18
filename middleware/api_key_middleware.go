// middleware/api_key_middleware.go
package middleware

import (
	"net/http"
	"os"
)

func APIKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		configuredApiKey := os.Getenv("API_KEY")
		if apiKey != configuredApiKey {
			http.Error(w, "Invalid API key", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
