package delete

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/store"
)

type Delete struct {
	Operations *ops.Operation
}

func InitDelete(o *ops.Operation) *Delete {
	return &Delete{
		Operations: o,
	}
}

func (d *Delete) NewDeleteRegistered(fileName string) error {
	parts := strings.Split(fileName, ".")
	if len(parts) < 2 {
		logger.GLogger.AddToLog("ERROR", "invalid file name for delete")
		return fmt.Errorf("invalid file name for delete")
	}

	nameParts := strings.Split(parts[0], "-")
	if len(nameParts) < 3 {
		return fmt.Errorf("invalid delete file format")
	}

	if nameParts[0] != "d" {
		return fmt.Errorf("not a delete file")
	}

	timeStr := nameParts[1]
	if len(timeStr) < 2 {
		return fmt.Errorf("invalid time format")
	}

	timeValue := timeStr[:len(timeStr)-1]
	timeUnit := timeStr[len(timeStr)-1:]

	timeToDeleteInt, err := strconv.Atoi(timeValue)
	if err != nil {
		return fmt.Errorf("invalid time value")
	}

	var deletionTime time.Time
	now := time.Now()

	switch timeUnit {
	case "d":
		deletionTime = now.Add(time.Duration(timeToDeleteInt) * 24 * time.Hour)
	case "h":
		deletionTime = now.Add(time.Duration(timeToDeleteInt) * time.Hour)
	case "m":
		deletionTime = now.Add(time.Duration(timeToDeleteInt) * time.Minute)
	case "s":
		deletionTime = now.Add(time.Duration(timeToDeleteInt) * time.Second)
	default:
		return fmt.Errorf("invalid time unit: must be d, h, m, or s")
	}
	id, err := store.GenerateStoreID(d.Operations.Store)
	if err != nil {
		return err
	}

	storeFile := store.StoredData{
		ID:   id,
		Task: "DELETE",
		Args: []string{
			fileName,
			deletionTime.String(),
		},
		InProgress: false,
	}

	d.Operations.Store.AddStoredData(storeFile)

	return nil
}
