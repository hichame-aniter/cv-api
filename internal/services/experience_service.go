package services

import (
	apperrors "cv-api/internal/errors"
	"cv-api/internal/models"
	"cv-api/internal/storage"
)

func GetExperiences() ([]models.Experience, error) {
	cv, err := storage.LoadCV()
	if err != nil {
		return nil, err
	}
	return cv.Work, nil
}
func GetExperience(id string) (*models.Experience, error) {
	cv, err := storage.LoadCV()
	if err != nil {
		return nil, err
	}

	for _, exp := range cv.Work {
		if exp.ID == id {
			return &exp, nil
		}
	}

	// return nil, fmt.Errorf("experience with id %s: %w", id, apperrors.ErrNotFound)
	return nil, apperrors.ErrNotFound
}
func UpdateExperience(id string, updatedExp models.Experience) error {
	cv, err := storage.LoadCV()
	if err != nil {
		return err
	}

	found := false
	for i, exp := range cv.Work {
		if exp.ID == id {
			cv.Work[i] = updatedExp
			found = true
			break
		}
	}

	if !found {
		return apperrors.ErrInvalid
	}

	return storage.SaveCV(cv)
}
