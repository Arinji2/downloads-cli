package watcher

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops/delete"
	"github.com/Arinji2/downloads-cli/ops/move"
	"github.com/Arinji2/downloads-cli/utils"
)

type DeletedFile struct {
	Path      string
	Timestamp time.Time
}

type WatcherLog struct {
	DeletedFiles []DeletedFile
	DeleteJobs   delete.Delete
	MoveJobs     move.Move
}

func (w *WatcherLog) FileCreated(path string) {
	filename := strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	logger.GLogger.Notify(fmt.Sprintf("Created File %s", filename))
	operationType, err := utils.GetOperationType(filename)
	if err != nil {
		logger.GLogger.AddToLog("ERROR", err.Error())
		return
	}
	switch operationType {
	case "d":
		err := w.DeleteJobs.NewDeleteRegistered(filename, path)
		if err != nil {
			err = fmt.Errorf("error creating delete job: %v", err)
			logger.GLogger.AddToLog("ERROR", err.Error())
		}
	case "md":
		err := w.MoveJobs.NewMoveRegistered(filename, path)
		if err != nil {
			err = fmt.Errorf("error creating move preset job: %v", err)
			logger.GLogger.AddToLog("ERROR", err.Error())
		}
	case "mc":
		err := w.MoveJobs.NewMoveRegistered(filename, path)
		if err != nil {
			err = fmt.Errorf("error creating move custom job: %v", err)
			logger.GLogger.AddToLog("ERROR", err.Error())
		}
	default:
		logger.GLogger.AddToLog("ERROR", fmt.Sprintf("invalid operation type: %s", operationType))
	}
}

func (w *WatcherLog) FileDeleted(path string, timestamp time.Time) {
	println("File deleted: " + path)
	w.DeletedFiles = append(w.DeletedFiles, DeletedFile{
		Path:      path,
		Timestamp: timestamp,
	})
}

func (w *WatcherLog) FileRenamed(path string, originalPath string, timestamp time.Time) {
	println("File renamed: " + path)

	// Find the index of the original file in DeletedFiles
	index := -1
	for i, df := range w.DeletedFiles {
		if df.Path == originalPath {
			// Check if the rename happened within a minute of deletion
			timeDiff := timestamp.Sub(df.Timestamp)
			if timeDiff <= time.Minute {
				index = i
				println("Found matching deleted file within 1 minute timeframe")
				break
			} else {
				println("Found matching deleted file but time difference was:", timeDiff.Seconds(), "seconds")
			}
		}
	}

	if index != -1 {
		w.DeletedFiles = slices.Delete(w.DeletedFiles, index, index+1)
		println("Removed Deleted File", path, "from WatcherLog")
	}
}
