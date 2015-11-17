package aggregators

import "github.com/apoydence/ledger/transaction"

func init() {
	AddToStore("count", NewCount())
}

func NewCount() AggregatorFunc {
	return AggregatorFunc(func(accounts []*transaction.Account) int64 {
		return int64(len(accounts))
	})
}
