package options

import (
	"encoding/json"
	"os"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/utils"
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
	utils.ChangeToGoModDir()
	_, err := os.Stat(OPTIONS_FILENAME)
	if err != nil || os.IsNotExist(err) {
		logger.GLogger.AddToLog("FATAL", "options file not found")
		os.Exit(1)
	}

	contents, err := os.ReadFile(OPTIONS_FILENAME)
	if err != nil {
		logger.GLogger.AddToLog("FATAL", err.Error())
		os.Exit(1)
	}

	var options Options
	err = json.Unmarshal(contents, &options)
	if err != nil {
		logger.GLogger.AddToLog("FATAL", err.Error())
		os.Exit(1)
	}

	return options
}
