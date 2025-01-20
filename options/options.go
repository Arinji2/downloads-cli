package options

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type CheckInterval struct {
	Delete int `json:"delete"`
	Move   int `json:"move"`
}
type Options struct {
	DownloadsFolder string            `json:"downloads_folder"`
	LogFile         string            `json:"log_file"`
	CheckInterval   CheckInterval     `json:"check_interval"`
	MovePresets     map[string]string `json:"move_presets"`
}

var OPTIONS_FILENAME = "options.json"

func GetOptions() Options {
	_, err := os.Stat(OPTIONS_FILENAME)
	if err != nil || os.IsNotExist(err) {
		files, err := os.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("could not find options.json, files in current directory:")
		for _, file := range files {
			log.Println(file.Name())
		}
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("current directory:", wd)
		os.Exit(1)
	}

	contents, err := os.ReadFile(OPTIONS_FILENAME)
	if err != nil {
		err = fmt.Errorf("failed to read options.json: %w", err)
		log.Fatal(err)
	}

	var options Options
	err = json.Unmarshal(contents, &options)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal options: %w", err)
		log.Fatal(err)
	}

	return options
}
