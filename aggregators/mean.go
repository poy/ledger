package aggregators

import (
	"github.com/poy/ledger/transaction"
)

func init() {
	AddToStore("mean", NewMean())
}

func NewMean() AggregatorFunc {
	return AggregatorFunc(func(accounts []*transaction.Account) string {
		var result transaction.Money
		for _, a := range accounts {
			result += a.Value
		}
		return (result / transaction.Money(len(accounts))).String()
	})
}
