package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/kairavkkp/dedoxify-backend/models"
)

type CreateDocumentRequest struct {
	UUID             string `json:"uuid"`
	FamilyID         uint   `json:"family_id"`
	MemberID         uint   `json:"member_id"`
	Category         string `json:"category"`
	IsThumbnailReady bool   `json:"is_thumbnail_ready"`
	IsProcessed      bool   `json:"is_processed"`
}

func CreateDocument(db DBInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db == nil {
			log.Println("Database interface is nil")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		var req CreateDocumentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid request payload"})
			return
		}

		document := models.Document{
			UUID:             uuid.MustParse(req.UUID),
			Category:         req.Category,
			IsThumbnailReady: req.IsThumbnailReady,
			IsProcessed:      req.IsProcessed,
			FamilyID:         req.FamilyID,
			MemberID:         req.MemberID,
		}

		result := db.Create(&document)
		if result.Error != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to create Document: " + result.Error.Error()})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(document)
	}
}
