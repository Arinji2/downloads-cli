package link

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops/core"
	"github.com/Arinji2/downloads-cli/store"
)

func FoundLink(data store.StoredData, l *Link) (bool, string, error) {
	core.UpdateProgress(data, true, l.Operations)
	defer l.Operations.Store.DeleteStoredData(data.ID)

	if l.Operations.IsTesting {
		return true, "", nil
	}
	linkType := LinkType(data.Args[1])
	path := filepath.Join(data.RelativePath)
	upload := new(Upload)
	upload.filePath = path
	switch linkType {
	case LinkTemp:
		upload.uploadType = LinkTemp
	case LinkPerm:
		upload.uploadType = LinkPerm
	default:
		return false, "", errors.New("invalid link type for switch")
	}
	d, err := upload.UploadData()
	if err != nil {
		return false, "", err
	}
	logger.GlobalLogger.AddToLog("INFO", fmt.Sprintf("Created link for file: %s", data.Args[0]))

	return true, d, nil
}
