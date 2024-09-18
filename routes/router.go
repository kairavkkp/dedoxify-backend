// routes/router.go
package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kairavkkp/dedoxify-backend/handlers"
)

// SetupRouter initializes and returns the Chi router
func SetupRouter() http.Handler {
	r := chi.NewRouter()

	// Define your routes
	r.Get("/", handlers.RootHandler) // Temporary root handler

	return r
}
