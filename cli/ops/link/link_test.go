package link_test

import (
	"testing"

	"github.com/Arinji2/downloads-cli/ops/link"
)

func TestNewLinkRegistered(t *testing.T) {
	s, tempDir, ops := setupTest(t)

	fileName, testFile := setupFS(t, tempDir, "test", link.LinkType("t"))
	linkJob := link.InitLink(ops, 0, "")

	if err := linkJob.NewLinkRegistered(fileName, testFile); err != nil {
		t.Fatalf("Failed to register new link: %v", err)
	}

	storedData, err := s.GetAllStoredData()
	if err != nil {
		t.Error(err)
	}
	if len(storedData) != 1 {
		t.Error("Expected 1 stored data, got ", len(storedData))
	}
}
