package main

import (
	"os"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/ops/delete"
	"github.com/Arinji2/downloads-cli/ops/link"
	"github.com/Arinji2/downloads-cli/ops/move"
	"github.com/Arinji2/downloads-cli/options"
	"github.com/Arinji2/downloads-cli/store"
	"github.com/Arinji2/downloads-cli/watcher"
)

func main() {
	opts := options.GetOptions()
	_, err := os.ReadDir(opts.DownloadsFolder)
	if err != nil {
		panic(err)
	}
	s := store.InitStore(true)
	log, err := logger.NewLogger(opts.LogFile, 1024*1024, "DOWNLOADS CLI")
	if err != nil {
		panic(err)
	}
	logger.GlobalizeLogger(log)
	deleteOps := ops.InitOperations("DELETE", s)
	deleteJob := delete.InitDelete(deleteOps, opts.CheckInterval.Delete)

	moveOps := ops.InitOperations("MOVE", s)
	moveJob := move.InitMove(moveOps, opts.CheckInterval.Move, opts.MovePresets)

	linkOps := ops.InitOperations("LINK", s)
	linkJob := link.InitLink(linkOps, opts.CheckInterval.Delete)

	go watcher.StartWatcher(s, opts, deleteJob, moveJob, linkJob)

	go deleteJob.RunDeleteJobs()
	go moveJob.RunMoveJobs()
	go linkJob.RunLinkJobs()

	log.Notify("DOWNLOADS CLI STARTED SUCCESSFULLY")
	log.AddToLog("INFO", "DOWNLOADS CLI STARTED SUCCESSFULLY")
	<-make(chan struct{})
}
