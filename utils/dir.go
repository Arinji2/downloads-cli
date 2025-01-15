package utils

import (
	"log"
	"os"
	"path/filepath"

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
