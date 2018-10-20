package filters

import (
	"fmt"

	"github.com/poy/ledger/database"
)

type FilterFactory interface {
	Generate(arg string) (database.Filter, error)
}

var store map[string]FilterFactory

func AddToStore(name string, filter FilterFactory) {
	if store == nil {
		store = make(map[string]FilterFactory)
	}

	if _, ok := store[name]; ok {
		panic("Duplicate filter names: " + name)
	}

	store[name] = filter
}

func Store() map[string]FilterFactory {
	return store
}

func Fetch(names ...string) ([]FilterFactory, error) {
	var results []FilterFactory
	for _, name := range names {
		filter, ok := store[name]
		if !ok {
			return nil, fmt.Errorf("Unknown filter name '%s'", name)
		}
		results = append(results, filter)
	}
	return results, nil
}
