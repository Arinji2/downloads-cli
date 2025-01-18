package move

import (
	"fmt"
	"strings"
	"time"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/store"
)

var (
	DEFAULT_MOVE_INTERVAL = 30
	CUSTOM_MOVE_SEPERATOR = "("
)

type MoveType string

const (
	MoveMD MoveType = "md"
	MoveMC MoveType = "mc"
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
	fmt.Println("FILENAME", fileName)
	destPath := CreateDestinationPath(strings.Split(fileName, "-")[1])
	if err != nil {
		return err
	}
	id, err := store.GenerateStoreID(m.Operations.Store)
	if err != nil {
		return err
	}
	fmt.Println("FILENAME", fileName)
	fmt.Println("PATHNAME", pathName)
	fmt.Println("DESTPATH", destPath)
	storeFile := store.StoredData{
		ID:   id,
		Task: "MOVE",
		Args: []string{
			fileName,
			pathName,
			destPath,
		},
		InProgress: false,
	}
	m.Operations.Store.AddStoredData(storeFile)

	return nil
}

func (m *Move) RunMoveJobs() {
	ticker := time.NewTicker(time.Second * time.Duration(m.CheckInterval))
	for range ticker.C {

		storedData, err := m.Operations.Store.GetAllStoredData()
		if err != nil {
			logger.GLogger.AddToLog("ERROR", err.Error())
			break
		}
		for _, data := range storedData {
			if data.Task == "MOVE" {
				if data.InProgress {
					continue
				}
				typeOfMove := MoveType(strings.Split(data.Args[0], "-")[0])
				switch typeOfMove {
				case MoveMD:
					_, err := FoundDefaultMove(data, m)
					if err != nil {
						err = fmt.Errorf("error handling default move job %v", err)
						logger.GLogger.AddToLog("ERROR", err.Error())
						continue
					}
				case MoveMC:
					_, err := FoundCustomMove(data, m)
					if err != nil {
						err = fmt.Errorf("error handling custom move job %v", err)
						logger.GLogger.AddToLog("ERROR", err.Error())
						continue
					}
				default:
					continue
				}
			}
		}
	}
}
