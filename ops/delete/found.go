package delete

import (
	"fmt"
	"os"
	"time"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/store"
)

func FoundDelete(data store.StoredData, d *Delete) {
	data.InProgress = true
	d.Operations.Store.UpdateStoredData(data.ID, data)
	currentTime := time.Now()
	layout := "2006-01-02 15:04:05.999999999 -0700 MST"
	deletionTime, err := time.Parse(layout, data.Args[1])
	if err != nil {
		logger.GLogger.AddToLog("ERROR", err.Error())
		d.Operations.Store.DeleteStoredData(data.ID)
		return
	}
	if currentTime.After(deletionTime) {

		err := os.Remove(data.Args[2])
		if err != nil {
			logger.GLogger.AddToLog("ERROR", err.Error())
			d.Operations.Store.DeleteStoredData(data.ID)
			return
		}
		d.Operations.Store.DeleteStoredData(data.ID)
		logger.GLogger.Notify(fmt.Sprintf("Finish Operation on File: %s", data.Args[0]))
		return
	} else {
		data.InProgress = false
		d.Operations.Store.UpdateStoredData(data.ID, data)
		return
	}
}
