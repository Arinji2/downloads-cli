package move

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops/core"
)

func verifyMove(fileName string, m *Move) (err error) {
	parts := strings.Split(fileName, ".")
	if len(parts) < 2 {
		logger.GlobalLogger.AddToLog("ERROR", "invalid file name for move")
		return fmt.Errorf("invalid file name for move")
	}

	nameParts := strings.Split(parts[0], "-")
	if len(nameParts) < 3 {
		return fmt.Errorf("invalid file name for move")
	}
	rawMoveType, err := core.GetOperationType(fileName)
	if err != nil {
		return err
	}
	moveStr := nameParts[1]
	moveType := MoveType(rawMoveType)
	switch moveType {
	case MoveMD:
		err = verifyMoveMD(m, moveStr)
		if err != nil {
			return err
		}
		return nil
	case MoveMC:
		err = verifyMoveMC(moveStr)
		if err != nil {
			return err
		}
		return nil
	case MoveMCD:

		locOfPrefix := strings.Index(moveStr, "#")
		if locOfPrefix == -1 {
			return fmt.Errorf("invalid move string for move default custom")
		}
		moveDefault := moveStr[:locOfPrefix]
		err = verifyMoveMD(m, moveDefault)
		if err != nil {
			return err
		}
		moveCustom := moveStr[locOfPrefix+1:]
		defaultPath := m.MovePresets[moveDefault]
		customPath := filepath.Join(defaultPath, moveCustom)
		err = verifyMoveMC(customPath)
		if err != nil {
			return err
		}
		return nil

	default:
		return fmt.Errorf("invalid move type")
	}
}

func verifyMoveMD(m *Move, moveStr string) error {
	if m.MovePresets[moveStr] == "" {
		return fmt.Errorf("invalid move string for move default")
	}
	return nil
}

func verifyMoveMC(moveStr string) error {
	destPath := CreateDestinationPath(moveStr)
	currentDir, _ := os.Getwd()
	destPath = filepath.Clean(destPath)
	checkCustomDirExists := os.Chdir(destPath)
	os.Chdir(currentDir)
	if checkCustomDirExists != nil {
		return fmt.Errorf("invalid move string for move custom")
	}
	return nil
}
