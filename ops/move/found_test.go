package move_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/ops/move"
	"github.com/Arinji2/downloads-cli/store"
)

func TestFoundDefaultMove(t *testing.T) {
	s := store.InitStore(true)
	ops := ops.InitTestingOperations("MOVE", s)
	moveJob := move.InitMove(ops, 0, map[string]string{
		"test": "/tmp/test",
	})

	err := moveJob.NewMoveRegistered("md-test-test1.txt", "/tmp/md-test-test1.txt")
	if err != nil {
		t.Error(err)
	}

	data, err := s.GetAllStoredData()
	if err != nil {
		t.Error(err)
	}
	if len(data) != 1 {
		t.Error("Expected 1 stored data, got ", len(data))
	}

	fileData := data[0]
	moved, err := move.FoundDefaultMove(fileData, moveJob)
	if err != nil {
		t.Error(err)
	}
	if !moved {
		t.Error("Expected moved to be true")
	}

	data, err = s.GetAllStoredData()
	if err != nil {
		t.Error(err)
	}
	if len(data) != 0 {
		t.Error("Expected 0 stored data, got ", len(data))
	}
}

func setupFS(t *testing.T, tempDir string, name string) (fileName, testFile, destPath string) {
	destPath = filepath.Join(tempDir, "test")
	os.Mkdir(destPath, 0755)
	formattedDestPath := strings.ReplaceAll(destPath, "/", "[")
	fileName = fmt.Sprintf("mc-%s-%s.txt", formattedDestPath, name)
	testFile = filepath.Join(tempDir, fileName)
	err := os.WriteFile(testFile, []byte("test"), 0644)
	if err != nil {
		t.Error(err)
	}
	return fileName, testFile, destPath
}

func TestFoundCustomMove_Valid(t *testing.T) {
	s := store.InitStore(true)
	ops := ops.InitTestingOperations("MOVE", s)
	tempDir := t.TempDir()
	moveJob := move.InitMove(ops, 0, map[string]string{})
	fileName, testFile, _ := setupFS(t, tempDir, "test")
	err := moveJob.NewMoveRegistered(fileName, testFile)
	if err != nil {
		t.Error(err)
	}
	data, err := s.GetAllStoredData()
	if err != nil {
		t.Error(err)
	}
	if len(data) != 1 {
		t.Error("Expected 1 stored data, got ", len(data))
	}

	fileData := data[0]
	moved, err := move.FoundCustomMove(fileData, moveJob)
	if err != nil {
		t.Error(err)
	}
	if !moved {
		t.Error("Expected moved to be true")
	}
}

func TestFoundCustomMove_Broken(t *testing.T) {
	s := store.InitStore(true)
	ops := ops.InitTestingOperations("MOVE", s)
	tempDir := t.TempDir()
	fileName, testFile, destPath := setupFS(t, tempDir, "brokenTest")
	moveJob := move.InitMove(ops, 0, map[string]string{})
	err := moveJob.NewMoveRegistered(fileName, testFile)
	if err != nil {
		t.Error(err)
	}
	data, err := s.GetAllStoredData()
	if err != nil {
		t.Error(err)
	}
	if len(data) != 1 {
		t.Error("Expected 1 stored data, got ", len(data))
	}

	fileData := data[0]
	fileData.Args[2] = fmt.Sprintf("%s/%s", destPath, "brokenTest")
	moved, err := move.FoundCustomMove(fileData, moveJob)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if moved {
		t.Error("Expected moved to be false")
	}

	storedData, err := s.GetAllStoredData()
	if err != nil {
		t.Error(err)
	}
	if len(storedData) != 0 {
		t.Error("Expected 0 stored data, got ", len(storedData))
	}
}
