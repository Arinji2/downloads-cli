package move_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Arinji2/downloads-cli/ops/move"
)

func validateRename(t *testing.T, destPath string) {
	_, err := os.Stat(destPath)
	if err != nil {
		if !os.IsNotExist(err) {
			t.Fatalf("Failed to stat file: %v", err)
		}
	}

	_, err = os.Stat(destPath)
	if err != nil {
		if !os.IsNotExist(err) {
			t.Fatalf("Failed to stat file: %v", err)
		}
	}
	dir := filepath.Dir(destPath)
	base := filepath.Base(destPath)
	parts := strings.Split(base, "-")
	fileName := parts[len(parts)-1]
	modified := filepath.Join(dir, fileName)

	modifiedFile, err := os.Stat(modified)
	if err != nil {
		t.Fatalf("Failed to stat file: %v", err)
	}
	if modifiedFile.IsDir() {
		t.Error("Expected file to be a file")
	}
}

func TestNewMoveRegistered_Default(t *testing.T) {
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
	destPath, err = moveJob.HandleMoveJob(storedData[0], move.MoveMD)
	if err != nil {
		t.Fatalf("Failed to handle move job: %v", err)
	}
	validateRename(t, destPath)
}
