package main

import (
	"cv-api/internal/config"
	"cv-api/internal/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorruptedCVLoad(t *testing.T) {
	router := routes.SetupRouter()
	assert.NotNil(t, router)

	config.DataPath = "data/corrupted_cv.json"

	// GetExperiences
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/basics", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	req, _ = http.NewRequest(http.MethodGet, "/experiences/0", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	config.LoadConfig() // reset config

}
