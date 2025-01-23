package options

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Arinji2/downloads-cli/ops/delete"
	"github.com/Arinji2/downloads-cli/ops/move"
	options_setup "github.com/Arinji2/downloads-cli/options/setup"
)

func (o *Options) SetupOptions(fileNotFound bool) *Options {
	if fileNotFound {
		options_setup.PrintMessage("Could not find options.json, starting options setup", "warn")
	} else {
		options_setup.PrintMessage("Found options.json, starting options setup", "info")
	}

	if o.DownloadsFolder == "" {
		o.DownloadsFolder = options_setup.InputDir("Enter the full path to your downloads folder:")
	}

	if o.LogFile == "" {
		file, err := options_setup.InputFileName("Enter the name for the log file (Type 1 to set as Default)", "Log", "1")
		if err != nil {
			o.LogFile = "app.log"
		} else {
			o.LogFile = fmt.Sprintf("%s.log", file)
		}
	}

	if o.CheckInterval.Delete == 0 {
		o.CheckInterval.Delete = delete.DEFAULT_DELETE_INTERVAL
		options_setup.PrintMessage(fmt.Sprintf("Checking every %d for deleted files. Update the options.json file to change.", delete.DEFAULT_DELETE_INTERVAL), "info")
	}

	if o.CheckInterval.Move == 0 {
		o.CheckInterval.Move = move.DEFAULT_MOVE_INTERVAL
		options_setup.PrintMessage(fmt.Sprintf("Checking every %d for moved files. Update the options.json file to change.", move.DEFAULT_MOVE_INTERVAL), "info")
	}

	initializeMovePresets(o)
	var file *os.File
	var err error
	if fileNotFound {
		file, err = os.Create(OPTIONS_FILENAME)
		if err != nil {
			options_setup.PrintMessage(err.Error(), "error")
		}

		options_setup.PrintMessage("Created options.json", "success")
	} else {
		file, err = os.OpenFile(OPTIONS_FILENAME, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			options_setup.PrintMessage(err.Error(), "error")
		}
	}
	timeStart := time.Now()
	defer file.Close()

	contents, err := json.Marshal(o)
	if err != nil {
		options_setup.PrintMessage(err.Error(), "error")
	}
	_, err = file.Write(contents)
	if err != nil {
		options_setup.PrintMessage(err.Error(), "error")
	}
	timeEnd := time.Now()
	options_setup.PrintMessage(fmt.Sprintf("Wrote to options.json in %s", timeEnd.Sub(timeStart)), "success")

	options_setup.PrintMessage("*********************************************************************************", "info")
	options_setup.PrintMessage("Setup Complete", "success")
	options_setup.PrintMessage("*********************************************************************************", "info")
	return o
}

func initializeMovePresets(o *Options) {
	options_setup.PrintMessage("Starting Default Move Presets Setup", "info")
	finishSetup := false

	for !finishSetup {
		movePreset, moveLocation := "", ""
		setup := false

		for !setup {
			options_setup.PrintMessage("Enter the name of the move preset (Example: picture). Type 0 to finish setup.", "info")
			fmt.Scanln(&movePreset)

			if options_setup.HandleExit(&finishSetup, movePreset) {
				break
			}

			if err := options_setup.ValidateMovePreset(movePreset, o.MovePresets); err != nil {
				options_setup.PrintMessage(err.Error(), "error")
				continue
			}

			moveLocation = options_setup.InputDir("Enter the full path to the move location (Example: /path/to/pictures):")
			if options_setup.HandleExit(&finishSetup, moveLocation) {
				break
			}

			if err := options_setup.ValidateMoveLocation(moveLocation, o.MovePresets); err != nil {
				options_setup.PrintMessage(err.Error(), "error")
				continue
			}

			o.MovePresets[movePreset] = moveLocation

			for k, v := range o.MovePresets {
				if k == movePreset {
					options_setup.PrintMessage(fmt.Sprintf("Preset %s -> %s **NEW**", k, v), "success")
				} else {
					options_setup.PrintMessage(fmt.Sprintf("Preset %s -> %s", k, v), "info")
				}
			}

			setup = options_setup.AskToAddAnotherLocation()
			if !setup {
				finishSetup = true
				break
			}
		}
	}
}
