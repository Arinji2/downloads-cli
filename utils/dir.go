package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Arinji2/downloads-cli/logger"
)

func ChangeToGoModDir() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			if err := os.Chdir(dir); err != nil {
				logger.GLogger.AddToLog("FATAL", err.Error())
				os.Exit(1)
			}
			return
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			logger.GLogger.AddToLog("FATAL", "Could not find go.mod file")
			os.Exit(1)
		}
		dir = parentDir
	}
}

func GetOperationType(fileName string) (string, error) {
	rawType := strings.Split(fileName, "-")[0]
	if rawType != "d" && rawType != "md" && rawType != "mc" {
		return "", fmt.Errorf("invalid operation type")
	}
	return rawType, nil
}

func WindowsMountIssue(inputString string) string {
	firstIndex := (strings.Index(inputString, ":")) + 1
	beforeMount := inputString[:firstIndex]
	afterMount := inputString[firstIndex:]
	println("1] BeforeMount", beforeMount)
	println("1] AfterMount", afterMount)
	afterMount = strings.ReplaceAll(afterMount, ":", "_")
	return beforeMount + afterMount
}
