package aggregators

import (
	"github.com/apoydence/ledger/transaction"
)

type AggregatorFunc func([]*transaction.Account) int64

func (a AggregatorFunc) Aggregate(accounts []*transaction.Account) int64 {
	return a(accounts)
}
