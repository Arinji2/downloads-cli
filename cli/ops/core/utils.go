package core

import (
	"strings"

	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/store"
)

func GetFilename(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

func UpdateProgress(data store.StoredData, progress bool, o *ops.Operation) {
	data.InProgress = progress
	o.Store.UpdateStoredData(data.ID, data)
}
