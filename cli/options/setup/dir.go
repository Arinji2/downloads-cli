package options_setup

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func InputDir(prompt string) string {
	for {
		PrintMessage(prompt, "info")
		var dir string
		fmt.Scanln(&dir)
		if err := checkDirPerms(dir); err != nil {
			PrintMessage(err.Error(), "error")
			wd, err := os.Getwd()
			if err != nil {
				PrintMessage(err.Error(), "error")
				continue
			}
			PrintMessage(fmt.Sprintf("Current Working Directory: %s", wd), "info")
			home, err := os.UserHomeDir()
			if err != nil {
				PrintMessage(err.Error(), "error")
				continue
			}
			PrintMessage(fmt.Sprintf("Home Directory: %s", home), "info")
			continue
		}
		files, err := os.ReadDir(dir)
		if err != nil {
			PrintMessage("No read permission in directory", "error")
			continue
		}
		PrintMessage(fmt.Sprintf("Directory set to: %s", dir), "success")
		PrintMessage(fmt.Sprintf("Found %d files in directory", len(files)), "info")
		return dir
	}
}

func ValidateMoveLocation(location string, existing map[string]string) error {
	if location == "" {
		return errors.New("path cannot be empty")
	}
	if err := checkDirPerms(location); err != nil {
		return err
	}
	for _, loc := range existing {
		if loc == location {
			return errors.New("location already exists")
		}
	}
	return nil
}

func checkDirPerms(dir string) error {
	info, err := os.Stat(dir)
	if err != nil {
		return errors.New("invalid directory")
	}
	if !info.IsDir() {
		return errors.New("not a directory")
	}
	testFile := filepath.Join(dir, ".write_test")
	if err := os.WriteFile(testFile, []byte("test"), 0666); err != nil {
		return errors.New("no write permission in directory")
	}
	os.Remove(testFile)
	return nil
}

func InputFileName(prompt, target, useDefault string) (string, error) {
	for {
		PrintMessage(prompt, "info")
		var fileName string
		fmt.Scanln(&fileName)
		if fileName == useDefault {
			return "", errors.New("using default name")
		}
		if fileName == "" {
			PrintMessage(fmt.Sprintf("%s name cannot be empty", target), "error")
			continue
		}
		if strings.Contains(fileName, ".") {
			PrintMessage(fmt.Sprintf("%s name cannot contain a dot", target), "error")
			continue
		}
		return fileName, nil
	}
}

func HandleExit(finishSetup *bool, input string) bool {
	if input == "0" {
		*finishSetup = true
		return true
	}
	return false
}
