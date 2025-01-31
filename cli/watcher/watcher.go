package watcher

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/options"
	"github.com/helshabini/fsbroker"
)

func StartWatcher(w *WatcherLog, opts *options.Options) {
	config := fsbroker.DefaultFSConfig()
	broker, err := fsbroker.NewFSBroker(config)
	if err != nil {
		logger.GlobalLogger.AddToLog("FATAL", fmt.Errorf("error creating FS Broker: %v", err).Error())
		log.Fatalf("error creating FS Broker: %v", err)
	}
	defer broker.Stop()

	if err := broker.AddRecursiveWatch(opts.DownloadsFolder); err != nil {
		logger.GlobalLogger.AddToLog("FATAL", fmt.Errorf("error adding watch: %v", err).Error())
		log.Fatalf("error adding watch: %v", err)
	}

	broker.Start()

	for {
		select {
		case event := <-broker.Next():
			fileRegex := `^[^-]+?-[^-]+?-[^-]+?\.[^.]+$`
			parts := strings.Split(event.Path, "/")
			c, err := regexp.Compile(fileRegex)
			if err != nil {
				fmt.Println(err)
			}
			match := c.MatchString(parts[len(parts)-1])
			if !match {
				continue
			}
			if event.Type.String() == "Create" {
				w.FileCreated(event.Path)
			}
			if event.Type.String() == "Remove" {
				w.FileDeleted(event.Path)
			}
			if event.Type.String() == "Rename" {
				originalPath, exists := event.Properties["OldPath"]
				if !exists {
					logger.GlobalLogger.AddToLog("ERROR", "originalPath not found in properties")
					continue
				}
				w.FileRenamed(event.Path, originalPath)
			}
		case error := <-broker.Error():
			logger.GlobalLogger.AddToLog("ERROR", error.Error())
			fmt.Println(error.Error())
		}
	}
}
