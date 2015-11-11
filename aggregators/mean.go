package aggregators

import (
	"github.com/apoydence/ledger/transaction"
)

func init() {
	AddToStore("mean", NewMean())
}

func NewMean() AggregatorFunc {
	return AggregatorFunc(func(accounts []*transaction.Account) float64 {
		var result float64
		for _, a := range accounts {
			result += a.Value
		}
		return result / float64(len(accounts))
	})
}
