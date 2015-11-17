package aggregators

import (
	"github.com/apoydence/ledger/transaction"
)

func init() {
	AddToStore("mean", NewMean())
}

func NewMean() AggregatorFunc {
	return AggregatorFunc(func(accounts []*transaction.Account) int64 {
		var result transaction.Money
		for _, a := range accounts {
			result += a.Value
		}
		return int64(result) / int64(len(accounts))
	})
}
