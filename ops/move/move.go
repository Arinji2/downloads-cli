package move

import (
	"time"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops"
	"github.com/Arinji2/downloads-cli/store"
)

var DEFAULT_MOVE_INTERVAL = 30

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
	moveType, moveStr, err := verifyMove(fileName)
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
			fileName,
			moveStr,
			pathName,
			moveType,
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
				// _, err := FoundDelete(data, m)
				// if err != nil {
				// 	logger.GLogger.AddToLog("ERROR", err.Error())
				// 	continue
				// }
			}
		}
	}
}
