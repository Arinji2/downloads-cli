package move

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Arinji2/downloads-cli/ops/core"
	"github.com/Arinji2/downloads-cli/store"
)

func FoundDefaultMove(data store.StoredData, m *Move) (moved bool, destPath string, err error) {
	core.UpdateProgress(data, true, m.Operations)

	originalPath := data.Args[0]
	destPath = m.MovePresets[data.Args[1]]
	fileName := core.GetFilename(originalPath)

	rawMoveType, err := core.GetOperationType(originalPath)
	if err != nil {
		return false, "", err
	}

	moveType := MoveType(rawMoveType)
	if moveType != MoveMD {
		return false, "", fmt.Errorf("invalid move type")
	}

	m.Operations.Store.DeleteStoredData(data.ID)
	if destPath == "" {
		return false, "", fmt.Errorf("invalid move string for move default")
	}

	moved, destPath, err = core.MoveFile(originalPath, destPath, fileName)
	if err != nil {
		return false, "", err
	}
	return moved, destPath, nil
}

func FoundCustomMove(data store.StoredData, m *Move) (moved bool, destPath string, err error) {
	data.InProgress = true
	m.Operations.Store.UpdateStoredData(data.ID, data)

	originalPath := data.Args[0]
	destPath = data.Args[1]
	fileName := core.GetFilename(originalPath)

	rawMoveType, err := core.GetOperationType(originalPath)
	if err != nil {
		return false, "", err
	}

	moveType := MoveType(rawMoveType)
	if moveType != MoveMC {
		return false, "", fmt.Errorf("invalid move type")
	}

	if !strings.HasSuffix(destPath, fileName) {
		destPath = filepath.Join(destPath, fileName)
	}

	m.Operations.Store.DeleteStoredData(data.ID)

	moved, _, err = core.MoveFile(originalPath, destPath, fileName)
	if err != nil {
		return false, "", err
	}

	return moved, destPath, nil
}

func FoundCustomDefaultMove(data store.StoredData, m *Move) (moved bool, destPath string, err error) {
	data.InProgress = true
	m.Operations.Store.UpdateStoredData(data.ID, data)

	originalPath := data.Args[0]
	destinationPath := data.Args[1]
	fileName := core.GetFilename(originalPath)

	rawMoveType, err := core.GetOperationType(originalPath)
	if err != nil {
		return false, "", err
	}

	moveType := MoveType(rawMoveType)
	if moveType != MoveMCD {
		return false, "", fmt.Errorf("invalid move type")
	}

	destPathParts := strings.Split(destinationPath, "/")
	destPath, ok := m.MovePresets[destPathParts[0]]
	if !ok {
		return false, "", fmt.Errorf("invalid move preset")
	}

	destPath = filepath.Join(destPath, filepath.Join(destPathParts[1:]...))
	m.Operations.Store.DeleteStoredData(data.ID)
	if destPath == "" {
		return false, "", fmt.Errorf("invalid move string for move default")
	}

	moved, _, err = core.MoveFile(originalPath, destPath, fileName)
	if err != nil {
		return false, "", err
	}

	return moved, destPath, nil
}

func renameToFilename(destPath string) (bool, error) {
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
