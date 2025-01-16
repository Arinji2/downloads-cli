package move

import (
	"fmt"
	"strings"

	"github.com/Arinji2/downloads-cli/logger"
)

func verifyMove(fileName string) (moveType string, moveStr string, err error) {
	parts := strings.Split(fileName, ".")
	if len(parts) < 2 {
		logger.GLogger.AddToLog("ERROR", "invalid file name for delete")
		return "", "", fmt.Errorf("invalid file name for delete")
	}

	nameParts := strings.Split(parts[0], "-")
	if len(nameParts) < 3 {
		return "", "", fmt.Errorf("invalid file name for delete")
	}
	moveType = nameParts[0]
	if moveType != "md" || moveType != "mc" {
		return "", "", fmt.Errorf("invalid file type for move")
	}

	moveStr = nameParts[1]
	if len(moveStr) == 0 {
		return "", "", fmt.Errorf("invalid move string for move")
	}
	return moveType, moveStr, nil
}
