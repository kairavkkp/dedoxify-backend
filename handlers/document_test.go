package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/kairavkkp/dedoxify-backend/db"
	"github.com/kairavkkp/dedoxify-backend/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func setupMockDocumentRouter(mockDB *db.MockDB) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/document", CreateDocument(mockDB))
	return r
}

func TestCreateDocument(t *testing.T) {
	mockDB := new(db.MockDB)

	document := models.Document{
		UUID:             uuid.New(),
		FamilyID:         1,
		MemberID:         1,
		Category:         "default",
		IsThumbnailReady: false,
		IsProcessed:      false,
		IsActive:         true,
	}

	mockDB.On("Create", mock.Anything).Return(&gorm.DB{})

	router := setupMockDocumentRouter(mockDB)

	body, _ := json.Marshal(document)
	req, _ := http.NewRequest("POST", "/document", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var result models.Document
	json.Unmarshal(resp.Body.Bytes(), &result)
	assert.Equal(t, document.UUID, result.UUID)
	mockDB.AssertExpectations(t)
}
