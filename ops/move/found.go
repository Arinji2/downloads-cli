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

func FoundDefaultMove(data store.StoredData, m *Move) (moved bool, err error) {
	data.InProgress = true
	m.Operations.Store.UpdateStoredData(data.ID, data)
	rawMoveType, err := utils.GetOperationType(data.Args[0])
	if err != nil {
		return false, err
	}
	moveType := MoveType(rawMoveType)
	if moveType != MoveMD {
		return false, fmt.Errorf("invalid move type")
	}
	fileName := data.Args[0]
	originalPath := data.Args[1]
	destPath := m.MovePresets[data.Args[2]]
	m.Operations.Store.DeleteStoredData(data.ID)
	if destPath == "" {
		return false, fmt.Errorf("invalid move string for move default")
	}
	if !strings.HasSuffix(destPath, fileName) {
		destPath = filepath.Join(destPath, fileName)
	}

	err = os.Rename(originalPath, destPath)
	if err != nil {
		return false, err
	}
	return true, nil
}

func FoundCustomMove(data store.StoredData, m *Move) (moved bool, err error) {
	data.InProgress = true
	m.Operations.Store.UpdateStoredData(data.ID, data)
	rawMoveType, err := utils.GetOperationType(data.Args[0])
	if err != nil {
		return false, err
	}
	moveType := MoveType(rawMoveType)
	if moveType != MoveMC {
		return false, fmt.Errorf("invalid move type")
	}
	fileName := data.Args[0]
	originalPath := data.Args[1]
	destPath := data.Args[2]
	fmt.Println("2] DESTPATH", destPath)
	fmt.Println("2] FILENAME", fileName)
	fmt.Println("2] HasSuffix", strings.HasSuffix(destPath, fileName))
	if !strings.HasSuffix(destPath, fileName) {
		destPath = filepath.Join(destPath, fileName)
	}
	m.Operations.Store.DeleteStoredData(data.ID)
	fmt.Println("3] DESTPATH", destPath)
	fmt.Println("3] FILENAME", fileName)
	fmt.Println("3] OriginalPath", originalPath)

	if runtime.GOOS == "windows" {
		firstIndex := strings.Index(destPath, ":")
		beforeMount := destPath[:firstIndex]
		afterMount := destPath[firstIndex:]
		println("1] BeforeMount", beforeMount)
		println("1] AfterMount", afterMount)
		afterMount = strings.ReplaceAll(afterMount, ":", "_")
		destPath = beforeMount + afterMount
	}
	err = os.Rename(originalPath, destPath)
	if err != nil {
		return false, err
	}
	return true, nil
}
