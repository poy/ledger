package aggregators

import (
	"github.com/apoydence/ledger/transaction"
)

func NewSum() AggregatorFunc {
	return AggregatorFunc(func(accounts []*transaction.Account) float64 {
		var result float64
		for _, a := range accounts {
			result += a.Value
		}
		return result
	})
}
