// routes/router.go
package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kairavkkp/dedoxify-backend/handlers"
	"gorm.io/gorm"
)

// SetupRouter initializes and returns the Chi router
func PublicRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()

	// r.Use(middleware.JWTAuth)

	r.Get("/", handlers.PublicRootHandler)
	r.Post("/v1/family", handlers.CreateFamily(db))
	r.Post("/v1/member", handlers.CreateMember(db))
	r.Post("/v1/document", handlers.CreateDocument(db))

	return r
}
