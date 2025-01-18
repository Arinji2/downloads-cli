package move_test

import (
	"testing"

	"github.com/Arinji2/downloads-cli/ops/move"
)

func TestNewMoveRegistered_Valid(t *testing.T) {
	s, tempDir, ops := setupTest(t)
	fileName, testFile, destPath := setupFS(t, tempDir, "md", "test")
	moveJob := move.InitMove(ops, 0, map[string]string{
		"test": destPath,
	})

	if err := moveJob.NewMoveRegistered(fileName, testFile); err != nil {
		t.Fatalf("Failed to register new move: %v", err)
	}

	storedData, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}
	if len(storedData) != 1 {
		t.Error("Expected 1 stored data, got ", len(storedData))
	}
}
