package link

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops/core"
	"github.com/Arinji2/downloads-cli/store"
)

func FoundLink(data store.StoredData, l *Link) (bool, string, error) {
	core.UpdateProgress(data, true, l.Operations)
	defer l.Operations.Store.DeleteStoredData(data.ID)

	linkType := LinkType(data.Args[1])
	path := data.RelativePath
	upload := new(Upload)
	upload.FilePath = path

	switch linkType {
	case LinkTemp:
		upload.UploadType = LinkTemp
	case LinkPerm:
		upload.UploadType = LinkPerm
	default:
		return false, "", errors.New("invalid link type for switch")
	}

	if l.UserHash != "" {
		upload.UserHash = l.UserHash
	}
	url, err := upload.UploadData()
	if err != nil {
		return false, "", err
	}
	logger.GlobalLogger.AddToLog("INFO", fmt.Sprintf("Created link for file: %s", data.Args[0]))

	lastIndex := strings.LastIndex(url, "/")
	urlID := url[lastIndex+1:]
	path = data.RelativePath
	linkedFile, linked, err := core.RenameToLink(urlID, string(linkType), path)
	if err != nil {
		logger.GlobalLogger.AddToLog("ERROR", err.Error())
		return false, "", err
	}

	if !linked {
		return false, "", fmt.Errorf("failed to rename file to link")
	}
	return true, linkedFile, nil
}
