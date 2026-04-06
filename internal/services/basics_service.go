package services

import (
	"cv-api/internal/config"
	"cv-api/internal/models"
	"cv-api/internal/storage"
)

func GetBasics() (*models.Basics, error) {
	cv, err := storage.LoadCV()
	if err != nil {
		return nil, err
	}

	return sanitizeBasics(cv.Basics), nil
}
func UpdateBasics(b models.Basics) error {
	cv, err := storage.LoadCV()
	if err != nil {
		return err
	}

	cv.Basics = b
	return storage.SaveCV(cv)
}
func sanitizeBasics(basics models.Basics) *models.Basics {
	// Create a copy to avoid modifying original
	sanitized := basics

	if config.Environment == "production" {
		// Hide sensitive data in production
		sanitized.Phone = "***" // or "***"
		sanitized.Location = models.Location{
			City:        basics.Location.City,
			Countrycode: basics.Location.Countrycode,
		}
		// Hide email domain or full email
		sanitized.Email = "***"
		sanitized.Url = "***"
		sanitized.Profiles = []models.Profile{}
	}

	return &sanitized
}
