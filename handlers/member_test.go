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

func setupMockMemberRouter(mockDB *db.MockDB) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/member", CreateMember(mockDB))
	return r
}

func TestCreateMember(t *testing.T) {
	mockDB := new(db.MockDB)

	Member := models.Member{
		UUID:      uuid.New(),
		FirstName: "Kairav",
		LastName:  "Pithadia",
		Email:     "kkp@gmail.com",
		FamilyID:  1,
	}

	mockDB.On("Create", mock.Anything).Return(&gorm.DB{})

	router := setupMockMemberRouter(mockDB)

	body, _ := json.Marshal(Member)
	req, _ := http.NewRequest("POST", "/member", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var result models.Member
	json.Unmarshal(resp.Body.Bytes(), &result)
	assert.Equal(t, Member.Email, result.Email)
	mockDB.AssertExpectations(t)
}
