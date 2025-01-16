package delete

import (
	"time"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/store"
)

var DEFAULT_DELETE_INTERVAL = 30

type Delete struct {
	Operations    *ops.Operation
	CheckInterval int
}

func InitDelete(o *ops.Operation, interval int) *Delete {
	if interval == 0 {
		interval = DEFAULT_DELETE_INTERVAL
	}
	return &Delete{
		Operations:    o,
		CheckInterval: interval,
	}
}

func (d *Delete) NewDeleteRegistered(fileName string, pathName string) error {
	timeStr, err := verifyDelete(fileName)
	if err != nil {
		return err
	}
	deletionTime, err := getDeletionTime(timeStr)
	if err != nil {
		return err
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
				_, err := FoundDelete(data, d)
				if err != nil {
					logger.GLogger.AddToLog("ERROR", err.Error())
					continue
				}

			}
		}
	}
}
