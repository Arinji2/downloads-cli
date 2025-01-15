package store_test

import (
	"os"
	"testing"

	"github.com/Arinji2/downloads-cli/store"
)

func TestAddingToStore(t *testing.T) {
	s := store.InitStore(true)
	s.AddStoredData(store.StoredData{ID: 1, Task: "test", Args: []string{"test"}, InProgress: false})

	file, err := os.Open(store.STORAGE_FILENAME)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	data, err := os.ReadFile(store.STORAGE_FILENAME)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	expected := "[{\"id\":1,\"task\":\"test\",\"args\":[\"test\"],\"in_progress\":false}]"
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}

func TestGettingStoredData(t *testing.T) {
	s := store.InitStore(true)
	s.AddStoredData(store.StoredData{ID: 1, Task: "test", Args: []string{"test"}, InProgress: false})

	storedData, foundData, err := s.GetStoredData(1)
	if err != nil {
		t.Errorf("Error getting stored data: %v", err)
	}

	if !foundData {
		t.Errorf("Expected data to be found")
	}

	if storedData.ID != 1 {
		t.Errorf("Expected ID 1, got %d", storedData.ID)
	}
	if storedData.Task != "test" {
		t.Errorf("Expected task test, got %s", storedData.Task)
	}
	if storedData.Args[0] != "test" {
		t.Errorf("Expected args test, got %s", storedData.Args[0])
	}
	if storedData.InProgress != false {
		t.Errorf("Expected in_progress false, got %t", storedData.InProgress)
	}
}

func TestGettingAllStoredData(t *testing.T) {
	s := store.InitStore(true)
	s.AddStoredData(store.StoredData{ID: 1, Task: "test", Args: []string{"test"}, InProgress: false})
	s.AddStoredData(store.StoredData{ID: 2, Task: "test2", Args: []string{"test2"}, InProgress: false})
	s.AddStoredData(store.StoredData{ID: 3, Task: "test3", Args: []string{"test3"}, InProgress: false})

	storedData, err := s.GetAllStoredData()
	if err != nil {
		t.Errorf("Error getting stored data: %v", err)
	}

	if len(storedData) != 3 {
		t.Errorf("Expected 3 stored data, got %d", len(storedData))
	}
}

func TestUpdatingStoredData(t *testing.T) {
	s := store.InitStore(true)
	s.AddStoredData(store.StoredData{ID: 1, Task: "test", Args: []string{"test"}, InProgress: false})

	s.UpdateStoredData(1, store.StoredData{ID: 1, Task: "test", Args: []string{"test2"}, InProgress: false})

	data, foundData, err := s.GetStoredData(1)
	if err != nil {
		t.Errorf("Error getting stored data: %v", err)
	}

	if !foundData {
		t.Errorf("Expected data to be found")
	}

	if data.ID != 1 {
		t.Errorf("Expected ID 1, got %d", data.ID)
	}
	if data.Task != "test" {
		t.Errorf("Expected task test, got %s", data.Task)
	}
	if data.Args[0] != "test2" {
		t.Errorf("Expected args test2, got %s", data.Args[0])
	}
	if data.InProgress != false {
		t.Errorf("Expected in_progress false, got %t", data.InProgress)
	}
}

func TestDeletingStoredData(t *testing.T) {
	s := store.InitStore(true)
	s.AddStoredData(store.StoredData{ID: 1, Task: "test", Args: []string{"test"}, InProgress: false})
	s.AddStoredData(store.StoredData{ID: 2, Task: "test2", Args: []string{"test2"}, InProgress: false})
	s.AddStoredData(store.StoredData{ID: 3, Task: "test3", Args: []string{"test3"}, InProgress: false})

	s.DeleteStoredData(1)

	storedData, err := s.GetAllStoredData()
	if err != nil {
		t.Errorf("Error getting stored data: %v", err)
	}

	if len(storedData) != 2 {
		t.Errorf("Expected 2 stored data, got %d", len(storedData))
	}

	_, foundData, err := s.GetStoredData(1)
	if err != nil {
		t.Errorf("Error getting stored data: %v", err)
	}

	if foundData {
		t.Errorf("Expected data to not be found")
	}
}

func TestClearingStore(t *testing.T) {
	s := store.InitStore(true)

	if err := s.AddStoredData(store.StoredData{ID: 1, Task: "test", Args: []string{"test"}, InProgress: true}); err != nil {
		t.Errorf("Error adding stored data: %v", err)
	}

	if err := s.ClearStore(); err != nil {
		t.Errorf("Error clearing store: %v", err)
	}

	storedData, err := s.GetAllStoredData()
	if err != nil {
		t.Errorf("Error getting stored data: %v", err)
	}

	if len(storedData) != 0 {
		t.Errorf("Expected 0 stored data, got %d", len(storedData))
	}
}
