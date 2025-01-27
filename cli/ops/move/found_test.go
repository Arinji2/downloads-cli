package move_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Arinji2/downloads-cli/ops/move"
)

func TestFoundDefaultMove_Valid(t *testing.T) {
	s, tempDir, ops := setupTest(t)
	fileName, testFile, destPath := setupFS(t, tempDir, "md", "test")
	moveJob := move.InitMove(ops, 0, map[string]string{
		"test": destPath,
	})

	if err := moveJob.NewMoveRegistered(fileName, testFile); err != nil {
		t.Fatalf("Failed to register new move: %v", err)
	}

	data, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}
	if len(data) != 1 {
		t.Fatalf("Expected 1 stored data, got %d", len(data))
	}
	moved, destPath, err := move.FoundDefaultMove(data[0], moveJob)
	if err != nil {
		t.Fatalf("FoundDefaultMove failed: %v", err)
	}
	if !moved {
		t.Error("Expected moved to be true")
	}
	if destPath == "" {
		t.Error("Expected destPath to be set")
	}
	file, err := os.Stat(destPath)
	if err != nil {
		t.Fatalf("Failed to stat file: %v", err)
	}
	if file.IsDir() {
		t.Error("Expected file to be a file")
	}

	data, err = s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}
	if len(data) != 0 {
		t.Errorf("Expected 0 stored data, got %d", len(data))
	}
}

func TestFoundDefaultMove_Broken(t *testing.T) {
	s, tempDir, ops := setupTest(t)
	fileName, testFile, destPath := setupFS(t, tempDir, "md", "test")

	moveJob := move.InitMove(ops, 0, map[string]string{
		"test": destPath,
	})

	if err := moveJob.NewMoveRegistered(fileName, testFile); err != nil {
		t.Fatalf("Failed to register new move: %v", err)
	}

	data, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}
	if len(data) != 1 {
		t.Fatalf("Expected 1 stored data, got %d", len(data))
	}

	data[0].Args[0] = "testBroken"
	moved, _, err := move.FoundDefaultMove(data[0], moveJob)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if moved {
		t.Error("Expected moved to be false")
	}

	data, err = s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}
	if len(data) != 0 {
		t.Errorf("Expected 0 stored data, got %d", len(data))
	}
}

func TestFoundCustomMove_Valid(t *testing.T) {
	s, tempDir, ops := setupTest(t)
	fileName, testFile, _ := setupFS(t, tempDir, "mc", "test")

	moveJob := move.InitMove(ops, 0, map[string]string{})
	if err := moveJob.NewMoveRegistered(fileName, testFile); err != nil {
		t.Fatalf("Failed to register new move: %v", err)
	}

	data, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}
	if len(data) != 1 {
		t.Fatalf("Expected 1 stored data, got %d", len(data))
	}

	moved, destPath, err := move.FoundCustomMove(data[0], moveJob)
	if err != nil {
		t.Fatalf("FoundCustomMove failed: %v", err)
	}
	if !moved {
		t.Error("Expected moved to be true")
	}

	if destPath == "" {
		t.Error("Expected destPath to be set")
	}

	file, err := os.Stat(destPath)
	if err != nil {
		t.Fatalf("Failed to stat file: %v", err)
	}
	if file.IsDir() {
		t.Error("Expected file to be a file")
	}
}

func TestFoundCustomMove_Broken(t *testing.T) {
	s, tempDir, ops := setupTest(t)
	fileName, testFile, destPath := setupFS(t, tempDir, "mc", "brokenTest")
	moveJob := move.InitMove(ops, 0, map[string]string{})
	if err := moveJob.NewMoveRegistered(fileName, testFile); err != nil {
		t.Fatalf("Failed to register new move: %v", err)
	}

	data, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}
	if len(data) != 1 {
		t.Fatalf("Expected 1 stored data, got %d", len(data))
	}

	data[0].Args[0] = filepath.Join(destPath, "brokenTest")
	moved, _, err := move.FoundCustomMove(data[0], moveJob)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if moved {
		t.Error("Expected moved to be false")
	}

	storedData, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}
	if len(storedData) != 0 {
		t.Errorf("Expected 0 stored data, got %d", len(storedData))
	}
}

func TestFoundCustomDefaultMove_Valid(t *testing.T) {
	s, tempDir, ops := setupTest(t)
	fileName, testFile, destPath := setupFS(t, tempDir, "mcd", "test")
	moveJob := move.InitMove(ops, 0, map[string]string{
		"default": destPath,
	})
	if err := moveJob.NewMoveRegistered(fileName, testFile); err != nil {
		t.Fatalf("Failed to register new move: %v", err)
	}

	data, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}

	if len(data) != 1 {
		t.Fatalf("Expected 1 stored data, got %d", len(data))
	}

	moved, destPath, err := move.FoundCustomDefaultMove(data[0], moveJob)
	if err != nil {
		t.Fatalf("FoundCustomDefaultMove failed: %v", err)
	}
	if !moved {
		t.Error("Expected moved to be true")
	}

	if destPath == "" {
		t.Error("Expected destPath to be set")
	}
	file, err := os.Stat(destPath)
	if err != nil {
		t.Fatalf("Failed to stat file: %v", err)
	}
	if file.IsDir() {
		t.Error("Expected file to be a file")
	}
}
