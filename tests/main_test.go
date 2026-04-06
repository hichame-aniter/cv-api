package main

import (
	"cv-api/internal/config"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func InitCVFile() {
	data, _ := os.ReadFile("data/standard_cv.json")
	_ = os.WriteFile("data/cv.json", data, 0644)
}

func TestMain(m *testing.M) {
	InitCVFile()
	config.LoadConfig()
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())

}
