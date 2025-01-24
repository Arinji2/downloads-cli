package move

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/utils"
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
	rawMoveType, err := utils.GetOperationType(fileName)
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
