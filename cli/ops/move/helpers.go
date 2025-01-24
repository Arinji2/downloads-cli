package move

import (
	"os"
	"path/filepath"
	"strings"
)

func CreateDestinationPath(rawPath string) string {
	var destPath string
	destPath = strings.ReplaceAll(rawPath, CUSTOM_MOVE_SEPERATOR, string(os.PathSeparator))

	// Handle home directory expansion
	if strings.HasPrefix(destPath, "~") {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			destPath = filepath.Join(homeDir, destPath[1:])
		}
	}
	return filepath.Clean(destPath)
}
