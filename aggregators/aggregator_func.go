package aggregators

import (
	"github.com/poy/ledger/transaction"
)

type AggregatorFunc func([]*transaction.Account) string

func (a AggregatorFunc) Aggregate(accounts []*transaction.Account) string {
	return a(accounts)
}
