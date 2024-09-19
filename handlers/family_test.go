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

func setupMockFamilyRouter(mockDB *db.MockDB) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/family", CreateFamily(mockDB))
	return r
}

func TestCreateFamily(t *testing.T) {
	mockDB := new(db.MockDB)

	family := models.Family{
		UUID:     uuid.New(),
		Name:     "Pithadia",
		IsActive: true,
	}

	mockDB.On("Create", mock.Anything).Return(&gorm.DB{})

	router := setupMockFamilyRouter(mockDB)

	body, _ := json.Marshal(family)
	req, _ := http.NewRequest("POST", "/family", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var result models.Family
	json.Unmarshal(resp.Body.Bytes(), &result)
	assert.Equal(t, family.Name, result.Name)
	mockDB.AssertExpectations(t)
}
