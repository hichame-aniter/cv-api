package storage

import (
	"cv-api/internal/config"
	"cv-api/internal/models"
	"encoding/json"
	"os"
	"sync"
)

func LoadCV() (*models.CV, error) {
	file, err := os.ReadFile(config.DataPath)
	if err != nil {
		return nil, err
	}

	var cv models.CV
	err = json.Unmarshal(file, &cv)
	if err != nil {
		return nil, err
	}

	return &cv, nil
}

var mu sync.Mutex

func SaveCV(cv *models.CV) error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.MarshalIndent(cv, "", " ")
	if err != nil {
		return err
	}

	tmp := config.DataPath + ".tmp"
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return err
	}

	return os.Rename(tmp, config.DataPath)
}
