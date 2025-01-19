package delete

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/utils"
)

func verifyDelete(fileName string) (timeStr string, err error) {
	parts := strings.Split(fileName, ".")
	if len(parts) < 2 {
		logger.GlobalLogger.AddToLog("ERROR", "invalid file name for delete")
		return "", fmt.Errorf("invalid file name for delete")
	}

	nameParts := strings.Split(parts[0], "-")
	if len(nameParts) < 3 {
		return "", fmt.Errorf("invalid file name for delete")
	}

	_, err = utils.GetOperationType(fileName)
	if err != nil {
		return "", err
	}
	timeStr = nameParts[1]
	if len(timeStr) < 2 {
		return "", fmt.Errorf("invalid file name for delete")
	}
	return timeStr, nil
}

func getDeletionTime(timeStr string) (time.Time, error) {
	timeValue := timeStr[:len(timeStr)-1]
	timeUnit := timeStr[len(timeStr)-1:]
	timeToDeleteInt, err := strconv.Atoi(timeValue)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time value")
	}

	var deletionTime time.Time
	now := time.Now()

	switch timeUnit {
	case "d":
		deletionTime = now.Add(time.Duration(timeToDeleteInt) * 24 * time.Hour)
	case "h":
		deletionTime = now.Add(time.Duration(timeToDeleteInt) * time.Hour)
	case "m":
		deletionTime = now.Add(time.Duration(timeToDeleteInt) * time.Minute)
	case "s":
		deletionTime = now.Add(time.Duration(timeToDeleteInt) * time.Second)
	default:
		return time.Time{}, fmt.Errorf("invalid time unit: must be d, h, m, or s")
	}
	return deletionTime, nil
}
