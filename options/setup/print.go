package options_setup

import (
	"fmt"
	"runtime"
)

func PrintMessage(message, messageType string) {
	var colorCode string
	switch runtime.GOOS {
	case "windows":
		colorCode = ""
	default:
		colorCode = map[string]string{
			"info":    "\033[34m", // Blue
			"warn":    "\033[33m", // Yellow
			"error":   "\033[31m", // Red
			"success": "\033[32m", // Green
		}[messageType]
	}
	fmt.Printf("%s%s\033[0m\n", colorCode, message)
}
