package main

import (
	"os"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/store"
)

var DOWNLOADS_FOLDER = "/home/arinji/Downloads/"

func main() {
	files, err := os.ReadDir(DOWNLOADS_FOLDER)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		println(file.Name())
	}
	store.InitStore(false)
	logger.InitLogger()
}
