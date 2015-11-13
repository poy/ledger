package aggregators

import "github.com/apoydence/ledger/transaction"

func init() {
	AddToStore("count", NewCount())
}

func NewCount() AggregatorFunc {
	return AggregatorFunc(func(accounts []*transaction.Account) float64 {
		return float64(len(accounts))
	})
}
