// routes/router.go
package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kairavkkp/dedoxify-backend/handlers"
	"github.com/kairavkkp/dedoxify-backend/middleware"
	"gorm.io/gorm"
)

// SetupRouter initializes and returns the Chi router
func PrivateRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.APIKeyAuth)
	r.Get("/", handlers.PrivateRootHandler)

	return r
}
