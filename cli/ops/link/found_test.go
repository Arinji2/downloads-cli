package link_test

import (
	"os"
	"testing"

	"github.com/Arinji2/downloads-cli/ops/link"
	_ "github.com/joho/godotenv/autoload"
)

func TestFoundLink_Perm(t *testing.T) {
	t.Parallel()
	s, tempDir, ops := setupTest(t)
	fileName, testFile := setupFS(t, tempDir, "test", link.LinkPerm)
	linkJob := link.InitLink(ops, 0, "")
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
		t.Fatalf("TestFoundLink_Perm failed: %v", err)
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

	deleteUploadedFile(t, destPath, userHash, link.LinkPerm)
}

func TestFoundLink_Temp(t *testing.T) {
	t.Parallel()
	typeOfLink := link.LinkTemp
	s, tempDir, ops := setupTest(t)
	fileName, testFile := setupFS(t, tempDir, "test", typeOfLink)
	linkJob := link.InitLink(ops, 0, "")

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
		t.Fatalf("TestFoundLink_Temp failed: %v", err)
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
}
