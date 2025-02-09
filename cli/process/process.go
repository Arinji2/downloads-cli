package process

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"syscall"

	"github.com/Arinji2/downloads-cli/logger"
)

// Retuns true if an instance of the program is already running
func PrerunProcessCheck() bool {
	statusFileExists := checkForStatusFile()
	if statusFileExists {
		statusText, err := os.ReadFile("status")
		if err != nil {
			os.Remove("status")
			return false
		}
		processID, err := strconv.Atoi(string(statusText))
		if err != nil {
			os.Remove("status")
			return false
		}
		exists := pidExists(int32(processID))
		if exists {
			return true
		} else {
			os.Remove("status")
			return false
		}
	}
	return false
}

// Adds the process ID to status file
func PostrunProcessCheck() bool {
	statusFileExists := checkForStatusFile()
	if statusFileExists {
		os.Remove("status")
	}

	statusFile, err := os.Create("status")
	if err != nil {
		err = fmt.Errorf("error creating status file: %v", err)
		logger.GlobalLogger.AddToLog("ERROR", err.Error())
		return false
	}
	defer statusFile.Close()
	currentID := os.Getpid()
	fmt.Fprintf(statusFile, "%d", currentID)

	return true
}

func checkForStatusFile() bool {
	_, err := os.Stat("status")
	return err == nil
}

// PidExists checks if a process with the given PID exists.
func pidExists(pid int32) bool {
	if pid <= 0 {
		return false
	}

	proc, err := os.FindProcess(int(pid))
	if err != nil {
		return false
	}

	err = proc.Signal(syscall.Signal(0))
	if err == nil {
		return true
	}

	if errors.Is(err, syscall.ESRCH) {
		return false
	} else if errors.Is(err, syscall.EPERM) {
		return true
	}

	return false
}
