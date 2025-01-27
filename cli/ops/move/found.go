package move

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Arinji2/downloads-cli/ops/core"
	"github.com/Arinji2/downloads-cli/store"
)

func FoundDefaultMove(data store.StoredData, m *Move) (moved bool, destPath string, err error) {
	core.UpdateProgress(data, true, m.Operations)
	defer m.Operations.Store.DeleteStoredData(data.ID)

	originalPath := data.Args[0]
	destPath = m.MovePresets[data.Args[1]]
	fileName := core.GetFilename(originalPath)

	if destPath == "" || originalPath == "" {
		return false, "", fmt.Errorf("invalid data for move default")
	}

	rawMoveType, err := core.GetOperationType(originalPath)
	if err != nil {
		return false, "", err
	}

	moveType := MoveType(rawMoveType)
	if moveType != MoveMD {
		return false, "", fmt.Errorf("invalid move type")
	}

	moved, destPath, err = core.MoveFile(originalPath, destPath, fileName)
	if err != nil {
		return false, "", err
	}
	return moved, destPath, nil
}

func FoundCustomMove(data store.StoredData, m *Move) (moved bool, destPath string, err error) {
	core.UpdateProgress(data, true, m.Operations)
	defer m.Operations.Store.DeleteStoredData(data.ID)

	originalPath := data.Args[0]
	destPath = data.Args[1]
	fileName := core.GetFilename(originalPath)

	if destPath == "" || originalPath == "" {
		return false, "", fmt.Errorf("invalid data for move default")
	}

	rawMoveType, err := core.GetOperationType(originalPath)
	if err != nil {
		return false, "", err
	}

	moveType := MoveType(rawMoveType)
	if moveType != MoveMC {
		return false, "", fmt.Errorf("invalid move type")
	}

	moved, destPath, err = core.MoveFile(originalPath, destPath, fileName)
	if err != nil {
		return false, "", err
	}

	return moved, destPath, nil
}

func FoundCustomDefaultMove(data store.StoredData, m *Move) (moved bool, destPath string, err error) {
	core.UpdateProgress(data, true, m.Operations)
	defer m.Operations.Store.DeleteStoredData(data.ID)

	originalPath := data.Args[0]
	destinationPath := data.Args[1]
	fileName := core.GetFilename(originalPath)

	if destinationPath == "" || originalPath == "" {
		return false, "", fmt.Errorf("invalid data for move default")
	}

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
	if destPath == "" {
		return false, "", fmt.Errorf("invalid move string for move default")
	}
	moved, destPath, err = core.MoveFile(originalPath, destPath, fileName)
	if err != nil {
		return false, "", err
	}

	return moved, destPath, nil
}
