package database

import "github.com/apoydence/ledger/transaction"

type Filter interface {
	Filter(*transaction.Transaction) []*transaction.Account
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
		if inTimeRange(t.Date, start, end) && len(filter(t, f)) > 0 {
			results = append(results, t)
		}
	}
	return results
}

func inTimeRange(date, start, end *transaction.Date) bool {
	return date.GreaterThanEqualTo(start) && end.GreaterThanEqualTo(date)
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
