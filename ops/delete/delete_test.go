package delete_test

import (
	"testing"

	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/ops/delete"
	"github.com/Arinji2/downloads-cli/store"
)

func TestNewDeleteRegistered(t *testing.T) {
	s := store.InitStore(true)
	ops := ops.InitTestingOperations("DELETE", s)
	deleteJob := delete.InitDelete(ops, 0)

	err := deleteJob.NewDeleteRegistered("d-10s-test1.txt", "/tmp/d-10s-test1.txt")
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
}
