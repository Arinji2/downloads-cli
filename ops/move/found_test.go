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
	"github.com/Arinji2/downloads-cli/utils"
)

func setupFS(t *testing.T, tempDir, moveType, name string) (fileName, testFile, destPath string) {
	switch moveType {
	case "md":
		destPath = "test"
	case "mc":
		destPath = filepath.Join(tempDir, "test")
	default:
		t.Error("Invalid move type")
	}
	os.Mkdir(destPath, 0755)
	formattedDestPath := strings.ReplaceAll(destPath, "/", "[")
	fileName = fmt.Sprintf("%s-%s-%s.txt", moveType, formattedDestPath, name)
	testFile = filepath.Join(tempDir, fileName)
	err := os.WriteFile(testFile, []byte("test"), 0644)
	if err != nil {
		t.Error(err)
	}
	return fileName, testFile, destPath
}

func TestFoundDefaultMove_Valid(t *testing.T) {
	s := store.InitStore(true)
	ops := ops.InitTestingOperations("MOVE", s)
	tempDir := t.TempDir()
	fileName, testFile, destPath := setupFS(t, tempDir, "md", "test")
	moveJob := move.InitMove(ops, 0, map[string]string{
		"test": destPath,
	})
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

	utils.ChangeToGoModDir()
	os.RemoveAll("/test")
}

func TestFoundDefaultMove_Broken(t *testing.T) {
	s := store.InitStore(true)
	ops := ops.InitTestingOperations("MOVE", s)
	tempDir := t.TempDir()
	fileName, testFile, destPath := setupFS(t, tempDir, "md", "test")
	moveJob := move.InitMove(ops, 0, map[string]string{
		"test": destPath,
	})
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
	fileData.Args[2] = "testBroken"
	moved, err := move.FoundDefaultMove(fileData, moveJob)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if moved {
		t.Error("Expected moved to be false")
	}

	data, err = s.GetAllStoredData()
	if err != nil {
		t.Error(err)
	}
	if len(data) != 0 {
		t.Error("Expected 0 stored data, got ", len(data))
	}

	utils.ChangeToGoModDir()
	wd, _ := os.Getwd()
	os.RemoveAll(filepath.Join(wd, "test"))
}

func TestFoundCustomMove_Valid(t *testing.T) {
	s := store.InitStore(true)
	ops := ops.InitTestingOperations("MOVE", s)
	tempDir := t.TempDir()
	moveJob := move.InitMove(ops, 0, map[string]string{})
	fileName, testFile, _ := setupFS(t, tempDir, "mc", "test")
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

	utils.ChangeToGoModDir()
	wd, _ := os.Getwd()
	os.RemoveAll(filepath.Join(wd, "test"))
}

func TestFoundCustomMove_Broken(t *testing.T) {
	s := store.InitStore(true)
	ops := ops.InitTestingOperations("MOVE", s)
	tempDir := t.TempDir()
	fileName, testFile, destPath := setupFS(t, tempDir, "mc", "brokenTest")
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
