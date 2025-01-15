package store

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/Arinji2/downloads-cli/logger"
)

func changeToGoModDir() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			if err := os.Chdir(dir); err != nil {
				logger.GLogger.AddToLog("FATAL", err.Error())
				os.Exit(1)
			}
			return
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			logger.GLogger.AddToLog("FATAL", "Could not find go.mod file")
			os.Exit(1)
		}
		dir = parentDir
	}
}

func readAndParseStoredData(s *Store) ([]StoredData, error) {
	if !s.cacheExpired {
		return s.cachedData, nil
	}

	data, err := os.ReadFile(STORAGE_FILENAME)
	if err != nil {
		logger.GLogger.AddToLog("ERROR", err.Error())
		return nil, err
	}

	var storedData []StoredData
	err = json.Unmarshal(data, &storedData)
	if err != nil {
		logger.GLogger.AddToLog("ERROR", err.Error())
		return nil, err
	}

	s.cachedData = storedData
	s.cacheExpired = false
	return storedData, nil
}

func generateStoreID(s *Store) (int, error) {
	storedData, err := readAndParseStoredData(s)
	if err != nil {
		return 0, err
	}
	var id int
	passes := 0
	for {
		passes++
		if passes > 1000 {
			return 0, fmt.Errorf("could not generate a unique ID")
		}
		id = rand.Intn(10000)
		exists := false
		for _, data := range storedData {
			if data.ID == id {
				exists = true
				break
			}
		}
		if !exists {
			break
		}
	}
	return id, nil
}
