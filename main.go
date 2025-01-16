package main

import (
	"os"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/options"
	"github.com/Arinji2/downloads-cli/store"
	"github.com/Arinji2/downloads-cli/watcher"
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

	watcher.StartWatcher(opts, s)
}
