package delete_test

import (
	"testing"

	"github.com/Arinji2/downloads-cli/ops/delete"
)

func TestNewDeleteRegistered(t *testing.T) {
	s, tempDir, ops := setupTest(t)

	fileName, testFile, _ := setupFS(t, tempDir, "test", "1d")
	deleteJob := delete.InitDelete(ops, 0)

	if err := deleteJob.NewDeleteRegistered(fileName, testFile); err != nil {
		t.Fatalf("Failed to register new move: %v", err)
	}

	storedData, err := s.GetAllStoredData()
	if err != nil {
		t.Error(err)
	}
	if len(storedData) != 1 {
		t.Error("Expected 1 stored data, got ", len(storedData))
	}
}
