package ops

import "github.com/Arinji2/downloads-cli/store"

type Operation struct {
	Name      string
	Store     *store.Store
	IsTesting bool
}

func InitOperations(name string, s *store.Store) *Operation {
	return &Operation{
		Name:      name,
		Store:     s,
		IsTesting: false,
	}
}

func InitTestingOperations(name string, s *store.Store) *Operation {
	return &Operation{
		Name:      name,
		Store:     s,
		IsTesting: true,
	}
}
