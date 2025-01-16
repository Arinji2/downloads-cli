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
	Operations    *ops.Operation
	CheckInterval int
}

func InitDelete(o *ops.Operation, interval int) *Delete {
	if interval == 0 {
		interval = 30
	}
	return &Delete{
		Operations:    o,
		CheckInterval: interval,
	}
}

func (d *Delete) NewDeleteRegistered(fileName string, pathName string) error {
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
	formattedTime := deletionTime.Format("2006-01-02 15:04:05.999999999 -0700 MST")
	storeFile := store.StoredData{
		ID:   id,
		Task: "DELETE",
		Args: []string{
			fileName,
			formattedTime,
			pathName,
		},
		InProgress: false,
	}

	d.Operations.Store.AddStoredData(storeFile)

	return nil
}

func (d *Delete) RunDeleteJobs() {
	ticker := time.NewTicker(time.Second * time.Duration(d.CheckInterval))
	for range ticker.C {

		storedData, err := d.Operations.Store.GetAllStoredData()
		if err != nil {
			logger.GLogger.AddToLog("ERROR", err.Error())
			break
		}
		for _, data := range storedData {
			if data.Task == "DELETE" {
				if data.InProgress {
					continue
				}
				FoundDelete(data, d)
			}
		}
	}
}
