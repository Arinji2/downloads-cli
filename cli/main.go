package main

import (
	"fmt"
	"os"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/ops/delete"
	"github.com/Arinji2/downloads-cli/ops/link"
	"github.com/Arinji2/downloads-cli/ops/move"
	"github.com/Arinji2/downloads-cli/options"
	"github.com/Arinji2/downloads-cli/process"
	"github.com/Arinji2/downloads-cli/store"
	"github.com/Arinji2/downloads-cli/watcher"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	opts := options.GetOptions()
	s := store.InitStore(false)

	log, err := logger.NewLogger(opts.LogFile, 1024*1024, "DOWNLOADS CLI")
	if err != nil {
		panic(err)
	}
	logger.GlobalizeLogger(log)

	exists := process.PrerunProcessCheck()
	if exists {
		log.Notify("Downloads CLI is already running")
		log.AddToLog("INFO", "Downloads CLI is already running")
		os.Exit(0)
	}

	setupOperations(s, &opts)

	log.Notify("DOWNLOADS CLI STARTED SUCCESSFULLY")
	log.AddToLog("INFO", "DOWNLOADS CLI STARTED SUCCESSFULLY")

	addedProcess := process.PostrunProcessCheck()
	if addedProcess {
		log.AddToLog("INFO", "Successfully added process to status file")
	} else {
		log.AddToLog("ERROR", "Failed to add process to status file")
		log.Notify("Failed to add process to status file")
		os.Exit(1)
	}

	channels := watcher.StartCLIWatcher()
	go func() {
		for range channels.UpdateOptions {
			fmt.Println("Received update signal")

			opts := options.GetOptions()
			setupOperations(s, &opts)
		}
	}()
	if <-channels.Exit {
		fmt.Println("Received exit signal, stopping program...")
		shutdown(s)
		os.Exit(0)
	}
}

func setupOperations(s *store.Store, o *options.Options) {
	deleteOps := ops.InitOperations("DELETE", s)
	deleteJob := delete.InitDelete(deleteOps, o.CheckInterval.Delete)

	moveOps := ops.InitOperations("MOVE", s)
	moveJob := move.InitMove(moveOps, o.CheckInterval.Move, o.MovePresets)

	linkOps := ops.InitOperations("LINK", s)
	var linkJob *link.Link
	if o.UserHash != "" {
		linkJob = link.InitLink(linkOps, o.CheckInterval.Delete, o.UserHash)
	} else {
		linkJob = link.InitLink(linkOps, o.CheckInterval.Delete, "")
	}

	w := watcher.WatcherLog{
		Store:      s,
		DeleteJobs: *deleteJob,
		MoveJobs:   *moveJob,
		LinkJobs:   *linkJob,
	}
	go watcher.StartWatcher(&w, o)

	go deleteJob.RunDeleteJobs()
	go moveJob.RunMoveJobs()
	go linkJob.RunLinkJobs()

	startup(o.DownloadsFolder, s, &w)
}
