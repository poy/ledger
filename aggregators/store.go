package aggregators

import (
	"fmt"
	"github.com/apoydence/ledger/database"
)

var store map[string]database.Aggregator

func AddToStore(name string, agg database.Aggregator) {
	if store == nil {
		store = make(map[string]database.Aggregator)
	}

	if _, ok := store[name]; ok {
		panic("Duplicate aggregator names: " + name)
	}

	store[name] = agg
}

func Store() map[string]database.Aggregator {
	return store
}

func Fetch(names ...string) ([]database.Aggregator, error) {
	var results []database.Aggregator
	for _, name := range names {
		agg, ok := store[name]
		if !ok {
			return nil, fmt.Errorf("Unknown aggregator name '%s'", name)
		}
		results = append(results, agg)
	}
	return results, nil
}
