package aggregators

import (
	"github.com/apoydence/ledger/transaction"
	"strconv"
)

func init() {
	AddToStore("count", NewCount())
}

func NewCount() AggregatorFunc {
	return AggregatorFunc(func(accounts []*transaction.Account) string {
		return strconv.Itoa(len(accounts))
	})
}
