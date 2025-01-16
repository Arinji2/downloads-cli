package watcher

import (
	"slices"
	"time"
)

type DeletedFile struct {
	Path      string
	Timestamp time.Time
}

type WatcherLog struct {
	DeletedFiles []DeletedFile
}

func FileCreated(path string) {
	println("File created: " + path)
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
