package move

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/ops/core"
	"github.com/Arinji2/downloads-cli/store"
)

var (
	DEFAULT_MOVE_INTERVAL = 5
	CUSTOM_MOVE_SEPERATOR = "#"
)

type MoveType string

const (
	MoveMD  MoveType = "md"
	MoveMC  MoveType = "mc"
	MoveMCD MoveType = "mcd"
)

type Move struct {
	Operations    *ops.Operation
	CheckInterval int
	MovePresets   map[string]string
}

func InitMove(o *ops.Operation, interval int, movePresets map[string]string) *Move {
	if interval == 0 {
		interval = DEFAULT_MOVE_INTERVAL
	}
	return &Move{
		Operations:    o,
		CheckInterval: interval,
		MovePresets:   movePresets,
	}
}

func (m *Move) NewMoveRegistered(fileName string, pathName string) error {
	err := verifyMove(fileName, m)
	destPath := CreateDestinationPath(strings.Split(fileName, "-")[1])
	if err != nil {
		return err
	}
	id, err := store.GenerateStoreID(m.Operations.Store)
	if err != nil {
		return err
	}
	storeFile := store.StoredData{
		ID:   id,
		Task: "MOVE",
		Args: []string{
			pathName,
			destPath,
		},
		InProgress: false,
	}
	m.Operations.Store.AddStoredData(storeFile)

	return nil
}

func (m *Move) HandleMoveJob(data store.StoredData, typeOfMove MoveType) (string, error) {
	switch typeOfMove {
	case MoveMD:
		moved, destPath, err := FoundDefaultMove(data, m)
		if err != nil {
			err = fmt.Errorf("error handling default move job %v", err)
			logger.GlobalLogger.AddToLog("ERROR", err.Error())
			return "", err
		}
		if !moved {
			return "", errors.New("default move job failed")
		}
		if _, err := core.RenameToFilename(destPath); err != nil {
			err = fmt.Errorf("error handling default move rename job %v", err)
			logger.GlobalLogger.AddToLog("ERROR", err.Error())
			return "", err
		}
		return destPath, nil
	case MoveMC:
		moved, destPath, err := FoundCustomMove(data, m)
		if err != nil {
			err = fmt.Errorf("error handling custom move job %v", err)
			logger.GlobalLogger.AddToLog("ERROR", err.Error())
			return "", err
		}
		if !moved {
			return "", errors.New("custom move job failed")
		}

		if _, err := core.RenameToFilename(destPath); err != nil {
			err = fmt.Errorf("error handling custom move rename job %v", err)
			logger.GlobalLogger.AddToLog("ERROR", err.Error())
			return "", err
		}

	case MoveMCD:
		moved, destPath, err := FoundCustomDefaultMove(data, m)
		if err != nil {
			err = fmt.Errorf("error handling custom default move job %v", err)
			logger.GlobalLogger.AddToLog("ERROR", err.Error())
			return "", err
		}
		if !moved {
			return "", errors.New("custom default move job failed")
		}

		if _, err := core.RenameToFilename(destPath); err != nil {
			err = fmt.Errorf("error handling custom move rename job %v", err)
			logger.GlobalLogger.AddToLog("ERROR", err.Error())
			return "", err
		}
	default:
		return "", errors.New("invalid move type")
	}
	return "", nil
}

func (m *Move) RunMoveJobs() {
	ticker := time.NewTicker(time.Second * time.Duration(m.CheckInterval))
	for range ticker.C {

		storedData, err := m.Operations.Store.GetAllStoredData()
		if err != nil {
			logger.GlobalLogger.AddToLog("ERROR", err.Error())
			break
		}
		for _, data := range storedData {
			if data.Task == "MOVE" {
				if data.InProgress {
					continue
				}
			}
			typeOfMove := MoveType(strings.Split(data.Args[0], "-")[0])
			_, err := m.HandleMoveJob(data, typeOfMove)
			if err != nil {
				logger.GlobalLogger.AddToLog("ERROR", err.Error())
				continue
			}

		}
	}
}
