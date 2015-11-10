package database

import "github.com/apoydence/ledger/transaction"

type Filter interface {
	Filter(*transaction.Transaction) bool
}

type Database struct {
	transactionList []*transaction.Transaction
}

func New() *Database {
	return new(Database)
}

func (db *Database) Add(ts ...*transaction.Transaction) {
	db.transactionList = append(db.transactionList, ts...)
}

func (db *Database) Query(start, end *transaction.Date, f Filter) []*transaction.Transaction {
	var results []*transaction.Transaction
	for _, t := range db.transactionList {
		if inTimeRange(t.Date, start, end) && filter(t, f) {
			results = append(results, t)
		}
	}
	return results
}

func inTimeRange(date, start, end *transaction.Date) bool {
	return date.GreaterThanEqualTo(start) && end.GreaterThanEqualTo(date)
}

func filter(t *transaction.Transaction, f Filter) bool {
	if f == nil {
		return true
	}

	return f.Filter(t)
}
