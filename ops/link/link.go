package link

import (
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
				_, err := FoundLink(data, l)
				if err != nil {
					logger.GlobalLogger.AddToLog("ERROR", err.Error())
					continue
				}

			}
		}
	}
}
