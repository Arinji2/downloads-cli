package ops

import "github.com/Arinji2/downloads-cli/store"

type Operation struct {
	Name    string
	Workers []chan store.StoredData
	Store   *store.Store
}

func InitOperations(name string, totalWorkers int, s *store.Store) *Operation {
	if totalWorkers == 0 {
		totalWorkers = 5
	}
	return &Operation{
		Name:    name,
		Workers: make([]chan store.StoredData, totalWorkers),
		Store:   s,
	}
}
