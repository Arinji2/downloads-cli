package watcher

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops/core"
	"github.com/Arinji2/downloads-cli/ops/delete"
	"github.com/Arinji2/downloads-cli/ops/link"
	"github.com/Arinji2/downloads-cli/ops/move"
	"github.com/Arinji2/downloads-cli/store"
)

type DeletedFile struct {
	Path      string
	Timestamp time.Time
}

type WatcherLog struct {
	Store      *store.Store
	DeleteJobs delete.Delete
	MoveJobs   move.Move
	LinkJobs   link.Link
}

func (w *WatcherLog) FileCreated(path string) bool {
	fileName := filepath.Base(path)
	operationType, err := core.GetOperationType(fileName)
	logger.GlobalLogger.AddToLog("INFO", fmt.Sprintf("operation type: %s", operationType))
	if err != nil {
		logger.GlobalLogger.AddToLog("ERROR", err.Error())
		return false
	}
	switch operationType {

	case "d":
		err := w.DeleteJobs.NewDeleteRegistered(fileName, path)
		if err != nil {
			err = fmt.Errorf("error creating delete job: %v", err)
			logger.GlobalLogger.AddToLog("ERROR", err.Error())
			return false
		}

	case "md":
		fallthrough
	case "mc":
		fallthrough
	case "mcd":
		err := w.MoveJobs.NewMoveRegistered(fileName, path)
		if err != nil {
			err = fmt.Errorf("error creating move job: %v", err)
			logger.GlobalLogger.AddToLog("ERROR", err.Error())
			return false
		}

	case "l":
		err := w.LinkJobs.NewLinkRegistered(fileName, path)
		if err != nil {
			err = fmt.Errorf("error creating link job: %v", err)
			logger.GlobalLogger.AddToLog("ERROR", err.Error())
			return false
		}

	default:
		logger.GlobalLogger.AddToLog("ERROR", fmt.Sprintf("invalid operation type: %s", operationType))
		return false
	}
	return true
}

func (w *WatcherLog) FileDeleted(path string) {
	parts := strings.Split(path, "/")
	filename := parts[len(parts)-1]
	err := w.DeleteJobs.DeleteByFilename(filename)
	if err != nil {
		logger.GlobalLogger.AddToLog("ERROR", err.Error())
		return
	}
}

func (w *WatcherLog) FileRenamed(path string, originalPath string) {
	originalParts := strings.Split(originalPath, "/")
	originalFilename := originalParts[len(originalParts)-1]

	newParts := strings.Split(path, "/")
	newFilename := newParts[len(newParts)-1]

	if newFilename == originalFilename {
		return
	}

	err := w.DeleteJobs.DeleteByFilename(originalFilename)
	if err != nil {
		logger.GlobalLogger.AddToLog("ERROR", err.Error())
	}
	w.FileCreated(path)
}
