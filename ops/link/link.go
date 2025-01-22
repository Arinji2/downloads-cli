package link

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/store"
)

var DEFAULT_LINK_INTERVAL = 5

type LinkType string

const (
	LinkTemp LinkType = "t"
	LinkPerm LinkType = "p"
)

type Link struct {
	Operations    *ops.Operation
	CheckInterval int
}

func InitLink(o *ops.Operation, interval int) *Link {
	if interval == 0 {
		interval = DEFAULT_LINK_INTERVAL
	}
	return &Link{
		Operations:    o,
		CheckInterval: interval,
	}
}

func (l *Link) NewLinkRegistered(fileName string, pathName string) error {
	linkType, err := verifyLink(fileName)
	if err != nil {
		return err
	}
	id, err := store.GenerateStoreID(l.Operations.Store)
	if err != nil {
		return err
	}
	fileInfo, err := os.Stat(pathName)
	if err != nil {
		return err
	}
	sizeInMB := float64(fileInfo.Size()) / (1024 * 1024)
	if sizeInMB > 100 {

		logger.GlobalLogger.Notify(fmt.Sprintf("File size is too large for %s. Max Size is %d(mb). Current File Size is %2.f", fileName, 100, sizeInMB))
		return errors.New("file size is too large")
	}
	storeFile := store.StoredData{
		ID:   id,
		Task: "LINK",
		Args: []string{
			fileName,
			string(linkType),
			pathName,
		},
		InProgress: false,
	}

	l.Operations.Store.AddStoredData(storeFile)

	return nil
}

func (l *Link) RunLinkJobs() {
	ticker := time.NewTicker(time.Second * time.Duration(l.CheckInterval))
	for range ticker.C {

		storedData, err := l.Operations.Store.GetAllStoredData()
		if err != nil {
			logger.GlobalLogger.AddToLog("ERROR", err.Error())
			break
		}
		for _, data := range storedData {
			if data.Task == "LINK" {
				if data.InProgress {
					continue
				}
				created, url, err := FoundLink(data, l)
				lastIndex := strings.LastIndex(url, "/")
				urlID := url[lastIndex+1:]
				if err != nil {
					logger.GlobalLogger.AddToLog("ERROR", err.Error())
					continue
				}
				if !created {
					continue
				}
				typeOfLink := LinkType(data.Args[1])
				if typeOfLink != LinkPerm && typeOfLink != LinkTemp {
					continue
				}
				path := data.Args[2]
				_, err = RenameToLink(urlID, typeOfLink, path)
				if err != nil {
					logger.GlobalLogger.AddToLog("ERROR", err.Error())
					continue
				}
				logger.GlobalLogger.Notify(fmt.Sprintf("Created link for file: %s", data.Args[0]))

			}
		}
	}
}

func RenameToLink(urlID string, typeOfLink LinkType, path string) (bool, error) {
	fileName := filepath.Base(path)
	url := fmt.Sprintf("dos.arinji.com#urlID=%s&type=%s", urlID, typeOfLink)

	newPath := strings.ReplaceAll(path, fileName, url)
	err := os.Rename(path, newPath)
	if err != nil {
		return false, err
	}
	return true, nil
}
