package ops

import "github.com/Arinji2/downloads-cli/store"

type Operation struct {
	Name  string
	Store *store.Store
}

func InitOperations(name string, s *store.Store) *Operation {
	return &Operation{
		Name:  name,
		Store: s,
	}
}
