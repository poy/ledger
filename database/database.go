package database

import (
	"github.com/apoydence/ledger/transaction"
	"time"
)

type Aggregator interface {
	Aggregate(acc []*transaction.Account) string
}

type Filter interface {
	Filter(*transaction.Transaction) []*transaction.Account
}

type Database struct {
	transactionList []*transaction.Transaction
}

func (db *Database) Add(ts ...*transaction.Transaction) {
	db.transactionList = append(db.transactionList, ts...)
}

func (db *Database) Aggregate(start, end time.Time, f Filter, aggs ...Aggregator) ([]*transaction.Transaction, []string) {
	results, accs := db.subQuery(start, end, f)
	if len(accs) == 0 {
		return nil, nil
	}

	var aggResults []string
	for _, agg := range aggs {
		aggResults = append(aggResults, agg.Aggregate(accs))
	}

	return results, aggResults
}

func (db *Database) Query(start, end time.Time, f Filter) []*transaction.Transaction {
	results, _ := db.subQuery(start, end, f)
	return results
}

func (db *Database) subQuery(start, end time.Time, f Filter) ([]*transaction.Transaction, []*transaction.Account) {
	var ts []*transaction.Transaction
	var accsResults []*transaction.Account

	for _, t := range db.transactionList {
		if !inTimeRange(t.Date, start, end) {
			continue
		}
		accs := filter(t, f)
		if len(accs) > 0 {
			ts = append(ts, t)
			accsResults = append(accsResults, accs...)
		}
	}
	return ts, accsResults
}

func inTimeRange(date, start, end time.Time) bool {
	return start.Add(-24*time.Hour).Before(date) && date.Before(end.Add(24*time.Hour))
}

func filter(t *transaction.Transaction, f Filter) []*transaction.Account {
	if t.Accounts == nil {
		return nil
	}

	if f == nil {
		return t.Accounts.Accounts
	}

	return f.Filter(t)
}
