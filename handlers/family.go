package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/kairavkkp/dedoxify-backend/models"
)

type CreateFamilyRequest struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

func CreateFamily(db DBInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db == nil {
			log.Println("Database interface is nil")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		var req CreateFamilyRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid request payload"})
			return
		}

		family := models.Family{
			UUID: uuid.MustParse(req.UUID),
			Name: req.Name,
		}

		result := db.Create(&family)
		if result.Error != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to create family: " + result.Error.Error()})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(family)
	}
}
