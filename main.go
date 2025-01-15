package main

import (
	"os"

	"github.com/Arinji2/downloads-cli/logger"
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
	store.InitStore(false)
	logger.InitLogger(opts.LogFile)
}
