package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

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
	s := store.InitStore(false)

	log, err := logger.NewLogger(opts.LogFile, 1024*1024, "DOWNLOADS CLI")
	if err != nil {
		panic(err)
	}
	logger.GlobalizeLogger(log)

	setupOperations(s, &opts)

	log.Notify("DOWNLOADS CLI STARTED SUCCESSFULLY")
	log.AddToLog("INFO", "DOWNLOADS CLI STARTED SUCCESSFULLY")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	fmt.Println("Shutting down gracefully...")
	s.Shutdown()
}

func setupOperations(s *store.Store, o *options.Options) {
	deleteOps := ops.InitOperations("DELETE", s)
	deleteJob := delete.InitDelete(deleteOps, o.CheckInterval.Delete)

	moveOps := ops.InitOperations("MOVE", s)
	moveJob := move.InitMove(moveOps, o.CheckInterval.Move, o.MovePresets)

	linkOps := ops.InitOperations("LINK", s)
	linkJob := link.InitLink(linkOps, o.CheckInterval.Delete)

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
}
