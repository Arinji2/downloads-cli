package core

import (
	"fmt"
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

func RenameToFilename(destPath string) (bool, error) {
	parts := strings.Split(destPath, "-")
	fileName := parts[len(parts)-1]
	dir := filepath.Dir(destPath)
	modified := filepath.Join(dir, fileName)
	err := os.Rename(destPath, modified)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RenameToLink(urlID string, typeOfLink string, path string) (string, bool, error) {
	fileName := filepath.Base(path)
	url := fmt.Sprintf("dos.arinji.com#urlID=%s&type=%s", urlID, typeOfLink)

	newPath := strings.ReplaceAll(path, fileName, url)
	err := os.Rename(path, newPath)
	if err != nil {
		return "", false, err
	}

	return newPath, true, nil
}
