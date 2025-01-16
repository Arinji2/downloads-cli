package watcher

import (
	"fmt"
	"log"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops/delete"
	"github.com/Arinji2/downloads-cli/options"
	"github.com/helshabini/fsbroker"
)

func StartWatcher(opts options.Options, deleteJob *delete.Delete) {
	config := fsbroker.DefaultFSConfig()
	broker, err := fsbroker.NewFSBroker(config)
	if err != nil {
		logger.GLogger.AddToLog("FATAL", fmt.Errorf("error creating FS Broker: %v", err).Error())
		log.Fatalf("error creating FS Broker: %v", err)
	}
	defer broker.Stop()

	if err := broker.AddRecursiveWatch(opts.DownloadsFolder); err != nil {
		logger.GLogger.AddToLog("FATAL", fmt.Errorf("error adding watch: %v", err).Error())
		log.Fatalf("error adding watch: %v", err)
	}

	broker.Start()
	watcherLog := WatcherLog{
		DeleteJobs: *deleteJob,
	}

	for {
		select {
		case event := <-broker.Next():
			if event.Type.String() == "Create" {
				watcherLog.FileCreated(event.Path)
			}
			if event.Type.String() == "Remove" {
				watcherLog.FileDeleted(event.Path, event.Timestamp)
			}
			if event.Type.String() == "Rename" {
				originalPath, exists := event.Properties["OldPath"]
				if !exists {
					logger.GLogger.AddToLog("ERROR", "originalPath not found in properties")
					continue
				}
				watcherLog.FileRenamed(event.Path, originalPath, event.Timestamp)
			}
		case error := <-broker.Error():
			logger.GLogger.AddToLog("ERROR", error.Error())
			fmt.Println(error.Error())
		}
	}
}
