package move_test

import (
	"testing"

	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/ops/move"
	"github.com/Arinji2/downloads-cli/store"
)

func TestNewMoveRegistered(t *testing.T) {
	s := store.InitStore(true)
	ops := ops.InitTestingOperations("DELETE", s)
	moveJob := move.InitMove(ops, 0, map[string]string{
		"test": "/tmp/test",
	})

	err := moveJob.NewMoveRegistered("md-test-test1.txt", "/tmp/md-test-test1.txt")
	if err != nil {
		t.Error(err)
	}

	storedData, err := s.GetAllStoredData()
	if err != nil {
		t.Error(err)
	}
	if len(storedData) != 1 {
		t.Error("Expected 1 stored data, got ", len(storedData))
	}

	err = s.ClearStore()
	if err != nil {
		t.Error(err)
	}

	err = moveJob.NewMoveRegistered("md-brokenTest-test1.txt", "/tmp/md-brokenTest-test1.txt")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	storedData, err = s.GetAllStoredData()
	if err != nil {
		t.Error(err)
	}
	if len(storedData) != 0 {
		t.Error("Expected 0 stored data, got ", len(storedData))
	}
}
