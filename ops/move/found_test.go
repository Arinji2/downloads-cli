package move_test

import (
	"testing"

	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/ops/move"
	"github.com/Arinji2/downloads-cli/store"
)

func TestFoundDefaultMove(t *testing.T) {
	s := store.InitStore(true)
	ops := ops.InitTestingOperations("MOVE", s)
	moveJob := move.InitMove(ops, 0, map[string]string{
		"test": "/tmp/test",
	})

	err := moveJob.NewMoveRegistered("md-test-test1.txt", "/tmp/md-test-test1.txt")
	if err != nil {
		t.Error(err)
	}

	data, err := s.GetAllStoredData()
	if err != nil {
		t.Error(err)
	}
	if len(data) != 1 {
		t.Error("Expected 1 stored data, got ", len(data))
	}

	fileData := data[0]
	moved, err := move.FoundDefaultMove(fileData, moveJob)
	if err != nil {
		t.Error(err)
	}
	if !moved {
		t.Error("Expected moved to be true")
	}

	data, err = s.GetAllStoredData()
	if err != nil {
		t.Error(err)
	}
	if len(data) != 0 {
		t.Error("Expected 0 stored data, got ", len(data))
	}
}
