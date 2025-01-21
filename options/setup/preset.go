package options_setup

import (
	"errors"
	"fmt"
	"strings"
)

func ValidateMovePreset(preset string, existing map[string]string) error {
	if preset == "" {
		return errors.New("move preset cannot be empty")
	}
	if strings.Contains(preset, "-") {
		return errors.New("move preset cannot contain a dash")
	}
	if _, exists := existing[preset]; exists {
		return errors.New("move preset already exists")
	}
	return nil
}

func AskToAddAnotherLocation() bool {
	PrintMessage("Add another location? y/n", "info")
	var addLoc string
	fmt.Scanln(&addLoc)
	return strings.ToLower(addLoc) != "n"
}
