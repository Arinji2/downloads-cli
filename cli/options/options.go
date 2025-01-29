package options

import (
	"encoding/json"
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
	UserHash        string            `json:"user_hash"` // Userhash from catbox.moe for perm links
}

var OPTIONS_FILENAME = "options.json"

func GetOptions() Options {
	if _, err := os.Stat(OPTIONS_FILENAME); os.IsNotExist(err) {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Current directory:", wd)

		var options Options
		options.MovePresets = make(map[string]string)
		options.SetupOptions(true)
		return options
	}

	contents, err := os.ReadFile(OPTIONS_FILENAME)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", OPTIONS_FILENAME, err)
	}

	var options Options
	options.MovePresets = make(map[string]string)
	if err := json.Unmarshal(contents, &options); err != nil {
		log.Fatalf("Failed to unmarshal options: %v", err)
	}

	return options
}
