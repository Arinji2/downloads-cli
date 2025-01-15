package main

import (
	"os"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/ops/delete"
	"github.com/Arinji2/downloads-cli/options"
	"github.com/Arinji2/downloads-cli/store"
)

func main() {
	opts := options.GetOptions()
	files, err := os.ReadDir(opts.DownloadsFolder)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		println(file.Name())
	}
	s := store.InitStore(true)
	logger.InitLogger(opts.LogFile)

	deleteOps := ops.InitOperations("DELETE", 0, s)
	deleteJob := delete.InitDelete(deleteOps)
	deleteJob.NewDeleteRegistered("d-3d-image.png")
}
