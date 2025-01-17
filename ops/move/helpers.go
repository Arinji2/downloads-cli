package move

import (
	"fmt"
	"os"
	"strings"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/utils"
)

func verifyMove(fileName string, m *Move) (err error) {
	parts := strings.Split(fileName, ".")
	if len(parts) < 2 {
		logger.GLogger.AddToLog("ERROR", "invalid file name for move")
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
		fmt.Println("HERE", m.MovePresets[moveStr], moveStr, m.MovePresets)
		if m.MovePresets[moveStr] == "" {
			return fmt.Errorf("invalid move string for move default")
		}
		return nil
	case MoveMC:
		destPath := CreateDesttinationPath(moveStr)
		currentDir, _ := os.Getwd()
		checkCustomDirExists := os.Chdir(destPath)
		os.Chdir(currentDir)
		if checkCustomDirExists != nil {
			fmt.Println(checkCustomDirExists.Error())
			return fmt.Errorf("invalid move string for move custom")
		}
		return nil
	default:
		return fmt.Errorf("invalid move type")
	}
}

func CreateDesttinationPath(rawPath string) string {
	var destPath string
	destPath = strings.ReplaceAll(rawPath, "[", "/")
	homeDir, _ := os.UserHomeDir()
	destPath = strings.ReplaceAll(destPath, "~", homeDir)
	return destPath
}
