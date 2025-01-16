package delete_test

import (
	"testing"
	"time"

	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/ops/delete"
	"github.com/Arinji2/downloads-cli/store"
)

func TestFoundDelete(t *testing.T) {
	s := store.InitStore(true)
	ops := ops.InitTestingOperations("DELETE", s)
	deleteJob := delete.InitDelete(ops, 0)

	err := deleteJob.NewDeleteRegistered("d-1s-test2.txt", "/tmp/d-1s-test2.txt")
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

	data, err = s.GetAllStoredData()
	if err != nil {
		t.Error(err)
	}
	if len(data) != 1 {
		t.Error("Expected 1 stored data, got ", len(data))
	}

	fileData := data[0]
	deleted, err := delete.FoundDelete(fileData, deleteJob)
	if err != nil {
		t.Error(err)
	}

	if deleted {
		t.Error("Expected Deleted to be false, got true")
	}

	time.Sleep(time.Second * 1)
	deleted, err = delete.FoundDelete(fileData, deleteJob)
	if err != nil {
		t.Error(err)
	}

	if !deleted {
		t.Error("Expected Deleted to be true, got false")
	}
}
