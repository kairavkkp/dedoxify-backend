package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/" , nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(RootHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected Status OK")

	expectedResponse := "Dedoxify Backend!"
	assert.Equal(t, expectedResponse, rr.Body.String(), "Expected response body to be Dedoxify Backend.")
}