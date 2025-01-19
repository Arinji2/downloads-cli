package main

import (
	"fmt"
	"os"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/ops/delete"
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
	logger.InitLogger(opts.LogFile)

	deleteOps := ops.InitOperations("DELETE", s)
	deleteJob := delete.InitDelete(deleteOps, opts.CheckInterval.Delete)

	moveOps := ops.InitOperations("MOVE", s)
	moveJob := move.InitMove(moveOps, opts.CheckInterval.Move, opts.MovePresets)

	go watcher.StartWatcher(s, opts, deleteJob, moveJob)

	go deleteJob.RunDeleteJobs()
	go moveJob.RunMoveJobs()

	fmt.Println("DOWNLOADS CLI STARTED SUCCESSFULLY")
	<-make(chan struct{})
}
