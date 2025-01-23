package move

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Arinji2/downloads-cli/store"
	"github.com/Arinji2/downloads-cli/utils"
)

func FoundDefaultMove(data store.StoredData, m *Move) (moved bool, destPath string, err error) {
	data.InProgress = true
	m.Operations.Store.UpdateStoredData(data.ID, data)
	rawMoveType, err := utils.GetOperationType(data.Args[0])
	if err != nil {
		return false, "", err
	}
	moveType := MoveType(rawMoveType)
	if moveType != MoveMD {
		return false, "", fmt.Errorf("invalid move type")
	}
	fileName := data.Args[0]
	originalPath := data.Args[1]
	destPath = m.MovePresets[data.Args[2]]
	m.Operations.Store.DeleteStoredData(data.ID)
	if destPath == "" {
		return false, "", fmt.Errorf("invalid move string for move default")
	}
	if !strings.HasSuffix(destPath, fileName) {
		destPath = filepath.Join(destPath, fileName)
	}

	err = os.Rename(originalPath, destPath)
	if err != nil {
		return false, "", err
	}
	return true, destPath, nil
}

func FoundCustomMove(data store.StoredData, m *Move) (moved bool, destPath string, err error) {
	data.InProgress = true
	m.Operations.Store.UpdateStoredData(data.ID, data)
	rawMoveType, err := utils.GetOperationType(data.Args[0])
	if err != nil {
		return false, "", err
	}
	moveType := MoveType(rawMoveType)
	if moveType != MoveMC {
		return false, "", fmt.Errorf("invalid move type")
	}
	fileName := data.Args[0]
	originalPath := data.Args[1]
	destPath = data.Args[2]
	if !strings.HasSuffix(destPath, fileName) {
		destPath = filepath.Join(destPath, fileName)
	}
	m.Operations.Store.DeleteStoredData(data.ID)
	if runtime.GOOS == "windows" {
		originalPath = utils.WindowsMountIssue(originalPath)
		destPath = utils.WindowsMountIssue(destPath)
	}
	err = os.Rename(originalPath, destPath)
	if err != nil {
		return false, "", err
	}
	return true, destPath, nil
}

func RenameToFilename(destPath string) (bool, error) {
	parts := strings.Split(destPath, "-")
	fileName := parts[len(parts)-1]
	dir := filepath.Dir(destPath)
	modified := filepath.Join(dir, fileName)
	err := os.Rename(destPath, modified)
	if err != nil {
		return false, err
	}
	return true, nil
}
