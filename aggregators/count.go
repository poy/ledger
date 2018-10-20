package aggregators

import (
	"github.com/poy/ledger/transaction"
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
