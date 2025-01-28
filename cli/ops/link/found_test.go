package link_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Arinji2/downloads-cli/ops/link"
	_ "github.com/joho/godotenv/autoload"
)

func TestFoundLink(t *testing.T) {
	s, tempDir, ops := setupTest(t)
	fileName, testFile := setupFS(t, tempDir, "test", link.LinkPerm)
	linkJob := link.InitLink(ops, 0)
	userHash, found := os.LookupEnv("USERHASH")
	if !found {
		t.Fatalf("USERHASH not found")
	}
	linkJob.UserHash = userHash

	if err := linkJob.NewLinkRegistered(fileName, testFile); err != nil {
		t.Fatalf("Failed to register new link: %v", err)
	}

	data, err := s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}
	if len(data) != 1 {
		t.Fatalf("Expected 1 stored data, got %d", len(data))
	}
	linked, destPath, err := link.FoundLink(data[0], linkJob)
	if err != nil {
		t.Fatalf("FoundDefaultMove failed: %v", err)
	}

	if !linked {
		t.Error("Expected moved to be true")
	}
	if destPath == "" {
		t.Error("Expected destPath to be set")
	}
	file, err := os.Stat(destPath)
	if err != nil {
		t.Fatalf("Failed to stat file: %v", err)
	}
	if file.IsDir() {
		t.Error("Expected file to be a file")
	}

	data, err = s.GetAllStoredData()
	if err != nil {
		t.Fatalf("Failed to get stored data: %v", err)
	}
	if len(data) != 0 {
		t.Errorf("Expected 0 stored data, got %d", len(data))
	}
	// Deleting the uploaded file
	fileName = filepath.Base(destPath)
	startingIndex := strings.Index(fileName, "#")
	endingIndex := strings.Index(fileName, "&")
	urlID := fileName[startingIndex+1 : endingIndex]
	indexOfEqual := strings.Index(urlID, "=")
	urlID = urlID[indexOfEqual+1:]
	if len(urlID) == 0 {
		t.Fatalf("Invalid urlID: %s", urlID)
	}

	upload := link.Upload{
		UploadType: link.LinkPerm,
		UserHash:   userHash,
	}
	err = upload.DeletedUploadedFile(urlID)
	if err != nil {
		t.Fatalf("Failed to delete uploaded file: %v", err)
	}
}
