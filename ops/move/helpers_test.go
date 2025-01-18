package move_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops"
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
	logger.InitLogger("")
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
