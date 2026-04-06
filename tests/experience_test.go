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

func TestGetExperiences(t *testing.T) {
	router := routes.SetupRouter()
	assert.NotNil(t, router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/experiences", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response []models.Experience
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
}
func TestGetExperience(t *testing.T) {
	router := routes.SetupRouter()
	assert.NotNil(t, router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/experiences/0", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response models.Experience
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
}
func TestGetExperienceNotFound(t *testing.T) {
	router := routes.SetupRouter()
	assert.NotNil(t, router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/experiences/-1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response models.Experience
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
}

func TestUpdateExperience(t *testing.T) {
	router := routes.SetupRouter()
	assert.NotNil(t, router)

	Experiences := models.Experience{}
	body, _ := json.Marshal(Experiences)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/experiences/0", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestUpdateExperienceBadRequest(t *testing.T) {
	router := routes.SetupRouter()
	assert.NotNil(t, router)

	Experiences := models.Experience{}
	body, _ := json.Marshal(Experiences)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/experiences/-1", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestUpdateExperienceInternalError(t *testing.T) {
	router := routes.SetupRouter()
	assert.NotNil(t, router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/experiences/0", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestExperienceConfigError(t *testing.T) {
	router := routes.SetupRouter()
	assert.NotNil(t, router)

	config.DataPath = "incorrectDataPath.json"

	// GetExperiences
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/experiences", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// UpdateExperiences
	w = httptest.NewRecorder()
	Experiences := models.Experience{}
	body, _ := json.Marshal(Experiences)
	req, _ = http.NewRequest(http.MethodPut, "/experiences/0", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	config.LoadConfig() // reset config

}
