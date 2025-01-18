package move_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/ops/move"
	"github.com/Arinji2/downloads-cli/store"
)

// setupFS creates necessary test files and directories
func setupFS(t *testing.T, tempDir, moveType, name string) (fileName, testFile, destPath string) {
	t.Helper()

	switch moveType {
	case "md":
		destPath = filepath.Join(tempDir, "test")
	case "mc":
		destPath = filepath.Join(tempDir, "test")
	default:
		t.Fatalf("Invalid move type: %s", moveType)
	}
	if err := os.MkdirAll(destPath, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	formattedDestPath := strings.ReplaceAll(destPath, "/", "[")
	if moveType == "md" {
		formattedDestPath = "test"
	}
	fileName = fmt.Sprintf("%s-%s-%s.txt", moveType, formattedDestPath, name)
	testFile = filepath.Join(tempDir, fileName)
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	return fileName, testFile, destPath
}

// setupTest initializes test environment with store and operations
func setupTest(t *testing.T) (*store.Store, string, *ops.Operation) {
	logger.InitLogger("log.txt")
	t.Helper()

	tempDir := t.TempDir()
	storeFile := filepath.Join(tempDir, "test-store.json")

	s := store.NewStore(storeFile)
	if err := s.Reset(); err != nil {
		t.Fatalf("Failed to init store: %v", err)
	}

	ops := ops.InitTestingOperations("MOVE", s)
	return s, tempDir, ops
}

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
	moved, err := move.FoundDefaultMove(data[0], moveJob)
	if err != nil {
		t.Fatalf("FoundDefaultMove failed: %v", err)
	}
	if !moved {
		t.Error("Expected moved to be true")
	}

	// Verify cleanup
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

	// Modify data to break it
	data[0].Args[2] = "testBroken"
	moved, err := move.FoundDefaultMove(data[0], moveJob)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if moved {
		t.Error("Expected moved to be false")
	}

	// Verify cleanup
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

	moved, err := move.FoundCustomMove(data[0], moveJob)
	if err != nil {
		t.Fatalf("FoundCustomMove failed: %v", err)
	}
	if !moved {
		t.Error("Expected moved to be true")
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

	data[0].Args[2] = filepath.Join(destPath, "brokenTest")
	moved, err := move.FoundCustomMove(data[0], moveJob)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if moved {
		t.Error("Expected moved to be false")
	}

	// Verify cleanup
	storedData, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}
	if len(storedData) != 0 {
		t.Errorf("Expected 0 stored data, got %d", len(storedData))
	}
}
