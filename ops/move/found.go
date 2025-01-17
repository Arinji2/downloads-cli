package move

import (
	"fmt"
	"os"
	"strings"

	"github.com/Arinji2/downloads-cli/store"
	"github.com/Arinji2/downloads-cli/utils"
)

func FoundDefaultMove(data store.StoredData, m *Move) (err error) {
	data.InProgress = true
	m.Operations.Store.UpdateStoredData(data.ID, data)
	rawMoveType, err := utils.GetOperationType(data.Args[0])
	if err != nil {
		return err
	}
	moveType := MoveType(rawMoveType)
	if moveType != MoveMD {
		return fmt.Errorf("invalid move type")
	}
	fileName := data.Args[0]
	originalPath := data.Args[1]
	destPath := m.MovePresets[data.Args[2]]
	if destPath == "" {
		return fmt.Errorf("invalid move string for move default")
	}
	if !strings.HasSuffix(destPath, fileName) {
		destPath = fmt.Sprintf("%s/%s", destPath, fileName)
	}

	m.Operations.Store.DeleteStoredData(data.ID)
	if m.Operations.IsTesting {
		return nil
	}
	err = os.Rename(originalPath, destPath)
	if err != nil {
		return err
	}
	return nil
}

func FoundCustomMove(data store.StoredData, m *Move) (err error) {
	data.InProgress = true
	m.Operations.Store.UpdateStoredData(data.ID, data)
	rawMoveType, err := utils.GetOperationType(data.Args[0])
	if err != nil {
		return err
	}
	moveType := MoveType(rawMoveType)
	if moveType != MoveMC {
		return fmt.Errorf("invalid move type")
	}
	fileName := data.Args[0]
	originalPath := data.Args[1]
	destPath := data.Args[2]
	if !strings.HasSuffix(destPath, fileName) {
		destPath = fmt.Sprintf("%s/%s", destPath, fileName)
	}

	m.Operations.Store.DeleteStoredData(data.ID)
	if m.Operations.IsTesting {
		return nil
	}
	err = os.Rename(originalPath, destPath)
	if err != nil {
		return err
	}
	return nil
}
