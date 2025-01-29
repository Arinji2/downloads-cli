package store_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/store"
)

func setupTest(t *testing.T) (*store.Store, string) {
	t.Helper()
	tempDir := t.TempDir()
	logger.SetupTestingLogger(t, tempDir)

	storeFile := filepath.Join(tempDir, "test-store.json")

	s := store.NewStore(storeFile)
	if err := s.Reset(); err != nil {
		t.Fatalf("Failed to init store: %v", err)
	}

	return s, tempDir
}

func TestAddingToStore(t *testing.T) {
	s, _ := setupTest(t)
	testData := store.StoredData{
		ID:         1,
		Task:       "test",
		Args:       []string{"test"},
		InProgress: false,
	}

	if err := s.AddStoredData(testData); err != nil {
		t.Fatalf("Failed to add data to store: %v", err)
	}

	// Verify file contents directly
	data, err := os.ReadFile(s.GetFilename())
	if err != nil {
		t.Fatalf("Error reading store file: %v", err)
	}

	var storedData []store.StoredData
	if err := json.Unmarshal(data, &storedData); err != nil {
		t.Fatalf("Error parsing store file: %v", err)
	}

	if len(storedData) != 1 {
		t.Errorf("Expected 1 item in store, got %d", len(storedData))
	}

	storeData := storedData[0]
	if storeData.ID != testData.ID {
		t.Errorf("Expected %v, got %v", testData, storedData[0])
	}
}

func TestGettingStoredData(t *testing.T) {
	s, _ := setupTest(t)
	testData := store.StoredData{
		ID:         1,
		Task:       "test",
		Args:       []string{"test"},
		InProgress: false,
	}

	if err := s.AddStoredData(testData); err != nil {
		t.Fatalf("Failed to add data to store: %v", err)
	}

	storedData, found, err := s.GetStoredData(1)
	if err != nil {
		t.Fatalf("Error getting stored data: %v", err)
	}

	if !found {
		t.Error("Expected data to be found")
	}

	if storedData.ID != testData.ID {
		t.Errorf("Expected %v, got %v", testData, storedData)
	}
}

func TestGettingAllStoredData(t *testing.T) {
	s, _ := setupTest(t)
	testData := []store.StoredData{
		{ID: 1, Task: "test1", Args: []string{"test1"}, InProgress: false},
		{ID: 2, Task: "test2", Args: []string{"test2"}, InProgress: false},
		{ID: 3, Task: "test3", Args: []string{"test3"}, InProgress: false},
	}

	for _, data := range testData {
		if err := s.AddStoredData(data); err != nil {
			t.Fatalf("Failed to add data to store: %v", err)
		}
	}

	storedData, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Error getting all stored data: %v", err)
	}

	if len(storedData) != len(testData) {
		t.Errorf("Expected %d stored items, got %d", len(testData), len(storedData))
	}

	for i, data := range storedData {
		if data.ID != testData[i].ID {
			t.Errorf("Item %d: expected %v, got %v", i, testData[i], data)
		}
	}
}

func TestUpdatingStoredData(t *testing.T) {
	s, _ := setupTest(t)
	originalData := store.StoredData{
		ID:         1,
		Task:       "test",
		Args:       []string{"test"},
		InProgress: false,
	}

	if err := s.AddStoredData(originalData); err != nil {
		t.Fatalf("Failed to add data to store: %v", err)
	}

	updatedData := store.StoredData{
		ID:         1,
		Task:       "test",
		Args:       []string{"test2"},
		InProgress: false,
	}

	if err := s.UpdateStoredData(1, updatedData); err != nil {
		t.Fatalf("Failed to update stored data: %v", err)
	}

	data, found, err := s.GetStoredData(1)
	if err != nil {
		t.Fatalf("Error getting stored data: %v", err)
	}

	if !found {
		t.Error("Expected data to be found after update")
	}

	if data.ID != updatedData.ID {
		t.Errorf("Expected %v, got %v", updatedData, data)
	}
}

func TestDeletingStoredData(t *testing.T) {
	s, _ := setupTest(t)
	testData := []store.StoredData{
		{ID: 1, Task: "test1", Args: []string{"test1"}, InProgress: false},
		{ID: 2, Task: "test2", Args: []string{"test2"}, InProgress: false},
		{ID: 3, Task: "test3", Args: []string{"test3"}, InProgress: false},
	}

	for _, data := range testData {
		if err := s.AddStoredData(data); err != nil {
			t.Fatalf("Failed to add data to store: %v", err)
		}
	}

	if err := s.DeleteStoredData(1); err != nil {
		t.Fatalf("Failed to delete stored data: %v", err)
	}

	storedData, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Error getting all stored data: %v", err)
	}

	if len(storedData) != 2 {
		t.Errorf("Expected 2 stored items after deletion, got %d", len(storedData))
	}

	_, found, err := s.GetStoredData(1)
	if err != nil {
		t.Fatalf("Error checking for deleted data: %v", err)
	}

	if found {
		t.Error("Expected deleted data to not be found")
	}
}

func TestClearingStore(t *testing.T) {
	s, _ := setupTest(t)
	testData := store.StoredData{
		ID:         1,
		Task:       "test",
		Args:       []string{"test"},
		InProgress: true,
	}

	if err := s.AddStoredData(testData); err != nil {
		t.Fatalf("Failed to add data to store: %v", err)
	}

	if err := s.ClearStore(); err != nil {
		t.Fatalf("Failed to clear store: %v", err)
	}

	storedData, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Error getting stored data after clear: %v", err)
	}

	if len(storedData) != 0 {
		t.Errorf("Expected empty store after clear, got %d items", len(storedData))
	}
}
