package watcher

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops/core"
	"github.com/Arinji2/downloads-cli/options"
	"github.com/Arinji2/downloads-cli/process"
	"github.com/helshabini/fsbroker"
)

func VerifyFile(path string) bool {
	fileRegex := `^[^-]+?-[^-]+?-[^-]+?\.[^.]+$`
	fileName := filepath.Base(path)
	c, err := regexp.Compile(fileRegex)
	if err != nil {
		fmt.Println(err)
	}
	match := c.MatchString(fileName)
	if !match {
		return false
	}

	_, err = core.GetOperationType(fileName)
	return err == nil
}

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
			verified := VerifyFile(event.Path)
			if !verified {
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

func StartStatusWatchr() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := fsbroker.DefaultFSConfig()
	broker, err := fsbroker.NewFSBroker(config)
	if err != nil {
		logger.GlobalLogger.AddToLog("FATAL", fmt.Errorf("error creating FS Broker For Status: %v", err).Error())
		log.Fatalf("error creating FS Broker: %v", err)
	}
	defer broker.Stop()

	if err := broker.AddWatch(currentDir); err != nil {
		logger.GlobalLogger.AddToLog("FATAL", fmt.Errorf("error adding watch to current directory for status: %v", err).Error())
		log.Fatalf("error adding watch: %v", err)
	}

	broker.Start()

	for {
		select {
		case event := <-broker.Next():
			fileName := filepath.Base(event.Path)
			if fileName != "status" {
				continue
			}
			if event.Type.String() == "Remove" {

				removed := process.EndProcessCheck()
				if !removed {
					continue
				}
				logger.GlobalLogger.AddToLog("INFO", "Status File Deleted, Exiting Process...")
				logger.GlobalLogger.Notify("Status File Deleted, Exiting Process...")
				os.Exit(0)
			}
		case error := <-broker.Error():
			err = fmt.Errorf("error in status watcher: %v", error)
			logger.GlobalLogger.AddToLog("ERROR", err.Error())
			fmt.Println(error.Error())
		}
	}
}
