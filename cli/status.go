package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/store"
	"github.com/Arinji2/downloads-cli/watcher"
)

func startup(downloadsFolder string, s *store.Store, w *watcher.WatcherLog) {
	count := 0
	data, err := s.GetAllStoredData()
	if err != nil {
		log.Fatalf("Error getting stored data: %s", err)
	}

	fileNames := make([]string, len(data))
	for i, d := range data {
		fileNames[i] = d.RelativePath
	}

	err = filepath.Walk(downloadsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == downloadsFolder {
			return nil
		}

		relPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}

		if !slices.Contains(fileNames, relPath) {
			added := w.FileCreated(relPath)
			if added {
				count++
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalf("Error walking through downloads folder: %s", err)
	}

	if count > 0 {
		logger.GlobalLogger.Notify(fmt.Sprintf("Added %d files to queue on startup", count))
	}
}

func shutdown(s *store.Store) {
	data, err := s.GetAllStoredData()
	if err != nil {
		err = fmt.Errorf("error getting all stored data for shutdown: %v", err)
		logger.GlobalLogger.AddToLog("ERROR", err.Error())
	}
	for i, v := range data {
		if v.InProgress {
			v.InProgress = false
			if err := s.UpdateStoredData(i, v); err != nil {
				err = fmt.Errorf("error updating stored data for shutdown: %v", err)
				logger.GlobalLogger.AddToLog("ERROR", err.Error())
			}
		}
	}
}
