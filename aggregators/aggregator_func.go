package aggregators

import (
	"github.com/apoydence/ledger/transaction"
)

type AggregatorFunc func([]*transaction.Account) float64

func (a AggregatorFunc) Aggregate(accounts []*transaction.Account) float64 {
	return a(accounts)
}
