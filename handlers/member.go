package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/kairavkkp/dedoxify-backend/models"
)

type CreateMemberRequest struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	FamilyID  uint   `json:"family_id"`
}

func CreateMember(db DBInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db == nil {
			log.Println("Database interface is nil")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		var req CreateMemberRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid request payload"})
			return
		}

		member := models.Member{
			UUID:      uuid.MustParse(req.UUID),
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			FamilyID:  req.FamilyID,
		}

		result := db.Create(&member)
		if result.Error != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to create Member: " + result.Error.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(member)
	}
}
