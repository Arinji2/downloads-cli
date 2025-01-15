package store

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/utils"
)

var STORAGE_FILENAME = "store.json"

type Store struct {
	storageFilename string
	cachedData      []StoredData
	cachedDataMutex sync.Mutex
	cacheExpired    bool
}
type StoredData struct {
	ID         int      `json:"id"`
	Task       string   `json:"task"`
	Args       []string `json:"args"`
	InProgress bool     `json:"in_progress"`
}

func InitStore(reset bool) *Store {
	utils.ChangeToGoModDir()
	_, err := os.Stat(STORAGE_FILENAME)
	if err != nil || os.IsNotExist(err) || reset {
		file, err := os.Create(STORAGE_FILENAME)
		if err != nil {
			logger.GLogger.AddToLog("FATAL", err.Error())
			logger.GLogger.Notify("Fatal Error in InitStore")
			os.Exit(1)
		}
		if _, err := file.WriteString("[]"); err != nil {
			logger.GLogger.AddToLog("FATAL", err.Error())
			logger.GLogger.Notify("Fatal Error in InitStore")
			os.Exit(1)
		}
		defer file.Close()
	}

	store := &Store{
		storageFilename: STORAGE_FILENAME,
		cachedData:      make([]StoredData, 0),
		cacheExpired:    true,
		cachedDataMutex: sync.Mutex{},
	}

	return store
}

func (s *Store) GetStoredData(id int) (StoredData, bool, error) {
	s.cachedDataMutex.Lock()
	defer s.cachedDataMutex.Unlock()

	if s.cacheExpired {
		if _, err := readAndParseStoredData(s); err != nil {
			logger.GLogger.AddToLog("ERROR", err.Error())
			return StoredData{}, false, err
		}
	}

	for _, data := range s.cachedData {
		if data.ID == id {
			return data, true, nil
		}
	}
	return StoredData{}, false, nil
}

func (s *Store) GetAllStoredData() ([]StoredData, error) {
	s.cachedDataMutex.Lock()
	defer s.cachedDataMutex.Unlock()

	if s.cacheExpired {
		if _, err := readAndParseStoredData(s); err != nil {
			logger.GLogger.AddToLog("ERROR", err.Error())
			return nil, err
		}
	}

	return s.cachedData, nil
}

func (s *Store) AddStoredData(data StoredData) error {
	s.cachedDataMutex.Lock()
	defer s.cachedDataMutex.Unlock()

	if s.cacheExpired {
		if _, err := readAndParseStoredData(s); err != nil {
			logger.GLogger.AddToLog("ERROR", err.Error())
			return err
		}
	}

	s.cachedData = append(s.cachedData, data)

	if err := s.saveToFile(); err != nil {
		logger.GLogger.AddToLog("ERROR", err.Error())
		return err
	}

	s.cacheExpired = false
	return nil
}

func (s *Store) UpdateStoredData(id int, data StoredData) error {
	s.cachedDataMutex.Lock()
	defer s.cachedDataMutex.Unlock()

	if s.cacheExpired {
		if _, err := readAndParseStoredData(s); err != nil {
			logger.GLogger.AddToLog("ERROR", err.Error())
			return err
		}
	}

	idFound := false
	for i, stored := range s.cachedData {
		if stored.ID == id {
			idFound = true
			s.cachedData[i] = data
			break
		}
	}

	if !idFound {
		logger.GLogger.AddToLog("ERROR", fmt.Sprintf("ID %d not found", id))
		return fmt.Errorf("id %d not found", id)
	}

	if err := s.saveToFile(); err != nil {
		logger.GLogger.AddToLog("ERROR", err.Error())
		return err
	}

	s.cacheExpired = false
	return nil
}

func (s *Store) DeleteStoredData(id int) error {
	s.cachedDataMutex.Lock()
	defer s.cachedDataMutex.Unlock()

	if s.cacheExpired {
		if _, err := readAndParseStoredData(s); err != nil {
			logger.GLogger.AddToLog("ERROR", err.Error())
			return err
		}
	}

	for i, stored := range s.cachedData {
		if stored.ID == id {
			s.cachedData = append(s.cachedData[:i], s.cachedData[i+1:]...)
			break
		}
	}

	if err := s.saveToFile(); err != nil {
		logger.GLogger.AddToLog("ERROR", err.Error())
		return err
	}

	s.cacheExpired = false
	return nil
}

func (s *Store) ClearStore() error {
	s.cachedDataMutex.Lock()
	defer s.cachedDataMutex.Unlock()

	newS := InitStore(true)
	s.cachedData = newS.cachedData
	s.cacheExpired = newS.cacheExpired
	s.cachedData = make([]StoredData, 0)
	s.cacheExpired = true
	return nil
}

func (s *Store) saveToFile() error {
	jsonData, err := json.Marshal(s.cachedData)
	if err != nil {
		logger.GLogger.AddToLog("ERROR", err.Error())
		return err
	}

	if err := os.WriteFile(s.storageFilename, jsonData, 0644); err != nil {
		logger.GLogger.AddToLog("ERROR", err.Error())
		return err
	}

	return nil
}
