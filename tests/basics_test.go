package main

import (
	"bytes"
	"cv-api/internal/config"
	"cv-api/internal/models"
	"cv-api/internal/routes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicsGetBasics(t *testing.T) {
	router := routes.SetupRouter()
	assert.NotNil(t, router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/basics", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response models.Basics
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", response.Name)
	assert.Equal(t, "Programmer", response.Label)
}

func TestUpdateBasics(t *testing.T) {
	router := routes.SetupRouter()
	assert.NotNil(t, router)

	basics := models.Basics{}
	body, _ := json.Marshal(basics)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/basics", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestUpdateBasicsError(t *testing.T) {
	router := routes.SetupRouter()
	assert.NotNil(t, router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/basics", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestConfigError(t *testing.T) {
	router := routes.SetupRouter()
	assert.NotNil(t, router)

	config.DataPath = "incorrectDataPath.json"

	// GetBasics
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/basics", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// UpdateBasics
	w = httptest.NewRecorder()
	basics := models.Basics{}
	body, _ := json.Marshal(basics)
	req, _ = http.NewRequest(http.MethodPut, "/basics", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	config.LoadConfig() // reset config
	// InitCVFile()
}
