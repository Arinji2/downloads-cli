package move

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
		if m.MovePresets[moveStr] == "" {
			return fmt.Errorf("invalid move string for move default")
		}
		return nil
	case MoveMC:
		destPath := CreateDestinationPath(moveStr)
		currentDir, _ := os.Getwd()
		destPath = filepath.Clean(destPath)
		checkCustomDirExists := os.Chdir(destPath)
		os.Chdir(currentDir)
		if checkCustomDirExists != nil {
			return fmt.Errorf("invalid move string for move custom")
		}
		return nil
	default:
		return fmt.Errorf("invalid move type")
	}
}

func CreateDestinationPath(rawPath string) string {
	// Handle Windows drive letter if present (e.g., C:)
	var drivePrefix string
	if runtime.GOOS == "windows" && len(rawPath) > 1 && rawPath[1] == ':' {
		drivePrefix = rawPath[:3] // Preserve "C:\"
		rawPath = rawPath[3:]
	}

	// Convert the path separators
	destPath := strings.ReplaceAll(rawPath, CUSTOM_MOVE_SEPERATOR, string(os.PathSeparator))

	// Handle home directory
	if strings.HasPrefix(destPath, "~") {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			destPath = filepath.Join(homeDir, destPath[1:])
		}
	}

	// Recombine with drive letter if on Windows
	if drivePrefix != "" {
		destPath = drivePrefix + destPath
	}

	return filepath.Clean(destPath)
}
