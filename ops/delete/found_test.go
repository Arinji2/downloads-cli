package delete_test

import (
	"testing"
	"time"

	"github.com/Arinji2/downloads-cli/ops/delete"
)

func TestFoundDelete_Valid(t *testing.T) {
	s, tempDir, ops := setupTest(t)

	fileName, testFile, _ := setupFS(t, tempDir, "test", "1s")
	deleteJob := delete.InitDelete(ops, 0)

	if err := deleteJob.NewDeleteRegistered(fileName, testFile); err != nil {
		t.Fatalf("Failed to register new move: %v", err)
	}

	data, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}
	if len(data) != 1 {
		t.Fatalf("Expected 1 stored data, got %d", len(data))
	}

	deleted, err := delete.FoundDelete(data[0], deleteJob)
	if err != nil {
		t.Fatalf("FoundDelete failed: %v", err)
	}
	if deleted {
		t.Error("Expected Deleted to be false")
	}

	time.Sleep(time.Second * 1)
	deleted, err = delete.FoundDelete(data[0], deleteJob)
	if err != nil {
		t.Error(err)
	}

	if !deleted {
		t.Error("Expected Deleted to be true")
	}
}

func TestFoundDelete_Invalid(t *testing.T) {
	s, tempDir, ops := setupTest(t)

	fileName, testFile, _ := setupFS(t, tempDir, "test", "2s")
	deleteJob := delete.InitDelete(ops, 0)

	if err := deleteJob.NewDeleteRegistered(fileName, testFile); err != nil {
		t.Fatalf("Failed to register new move: %v", err)
	}

	data, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}
	if len(data) != 1 {
		t.Fatalf("Expected 1 stored data, got %d", len(data))
	}

	deleted, err := delete.FoundDelete(data[0], deleteJob)
	if err != nil {
		t.Fatalf("FoundDelete failed: %v", err)
	}
	if deleted {
		t.Error("Expected Deleted to be false")
	}

	time.Sleep(time.Second * 1)
	deleted, err = delete.FoundDelete(data[0], deleteJob)
	if err != nil {
		t.Error(err)
	}

	if deleted {
		t.Error("Expected Deleted to be true")
	}
}
