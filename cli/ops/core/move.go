package core

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func MoveFile(originalPath, destPath, fileName string) (bool, string, error) {
	if !strings.HasSuffix(destPath, fileName) {
		destPath = filepath.Join(destPath, fileName)
	}

	if runtime.GOOS == "windows" {
		originalPath = WindowsMountIssue(originalPath)
		destPath = WindowsMountIssue(destPath)
	}

	err := os.Rename(originalPath, destPath)
	if err != nil {
		return false, "", err
	}

	return true, destPath, nil
}
