package store

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/Arinji2/downloads-cli/logger"
)

func readAndParseStoredData(s *Store) ([]StoredData, error) {
	if !s.cacheExpired {
		return s.cachedData, nil
	}

	data, err := os.ReadFile(s.storageFilename)
	if err != nil {
		logger.GlobalLogger.AddToLog("ERROR", err.Error())
		return nil, err
	}

	var storedData []StoredData
	err = json.Unmarshal(data, &storedData)
	if err != nil {
		storedData = make([]StoredData, 0)
	}

	s.cachedData = storedData
	s.cacheExpired = false
	return storedData, nil
}

func GenerateStoreID(s *Store) (int, error) {
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
