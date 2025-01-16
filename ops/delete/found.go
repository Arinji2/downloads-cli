package delete

import (
	"fmt"
	"os"
	"time"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/store"
)

func FoundDelete(data store.StoredData, d *Delete) (bool, error) {
	data.InProgress = true
	d.Operations.Store.UpdateStoredData(data.ID, data)

	currentTime := time.Now()
	layout := "2006-01-02 15:04:05.999999999 -0700 MST"
	deletionTime, err := time.Parse(layout, data.Args[1])
	if err != nil {
		d.Operations.Store.DeleteStoredData(data.ID)
		return false, err
	}

	if currentTime.After(deletionTime) {
		if d.Operations.IsTesting {
			return true, nil
		}
		err := os.Remove(data.Args[2])
		if err != nil {
			d.Operations.Store.DeleteStoredData(data.ID)
			return false, err
		}
		d.Operations.Store.DeleteStoredData(data.ID)
		logger.GLogger.AddToLog("INFO", fmt.Sprintf("Deleted file: %s", data.Args[0]))
		return true, nil
	} else {
		data.InProgress = false
		d.Operations.Store.UpdateStoredData(data.ID, data)
		return false, nil
	}
}
