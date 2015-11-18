package database

import (
	"github.com/apoydence/ledger/transaction"
	"time"
)

type Grouping struct {
	dbMap map[string]*Database
}

func NewGrouping() *Grouping {
	return &Grouping{
		dbMap: make(map[string]*Database),
	}
}

func (g *Grouping) Add(ts ...*transaction.Transaction) {
	for _, t := range ts {
		for _, acc := range t.Accounts.Accounts {
			db, ok := g.dbMap[acc.Name]
			if !ok {
				db = new(Database)
				g.dbMap[acc.Name] = db
			}
			db.Add(t)
		}
	}
}

func (g *Grouping) Aggregate(start, end time.Time, f Filter, aggs ...Aggregator) map[string][]string {
	result := make(map[string][]string)
	for accName, db := range g.dbMap {
		accNameFilter := newGroupFilter(accName)
		_, aggResults := db.Aggregate(start, end, CombineFilters(accNameFilter, f), aggs...)
		if len(aggResults) == 0 {
			continue
		}
		result[accName] = aggResults

	}
	return result
}

func newGroupFilter(name string) Filter {
	return FilterFunc(func(t *transaction.Transaction) []*transaction.Account {
		var results []*transaction.Account
		for _, acc := range t.Accounts.Accounts {
			if acc.Name == name {
				results = append(results, acc)
			}
		}
		return results
	})
}
