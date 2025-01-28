package link_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/ops/link"
	"github.com/Arinji2/downloads-cli/store"
)

// setupFS creates necessary test files and directories
func setupFS(t *testing.T, tempDir, name string, typeOfLink link.LinkType) (fileName, testFile string) {
	t.Helper()
	destPath := filepath.Join(tempDir, "test")
	if err := os.MkdirAll(destPath, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	fileName = fmt.Sprintf("l-%s-%s.txt", typeOfLink, name)
	testFile = filepath.Join(tempDir, fileName)
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	return fileName, testFile
}

// setupTest initializes test environment with store and operations
func setupTest(t *testing.T) (*store.Store, string, *ops.Operation) {
	log, err := logger.NewLogger("log.txt", 1024*1024, "TEST")
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	logger.GlobalizeLogger(log)

	t.Helper()

	tempDir := t.TempDir()
	storeFile := filepath.Join(tempDir, "test-store.json")

	s := store.NewStore(storeFile)
	if err := s.Reset(); err != nil {
		t.Fatalf("Failed to init store: %v", err)
	}

	ops := ops.InitTestingOperations("LINK", s)
	return s, tempDir, ops
}
