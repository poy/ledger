package filters

import (
	"github.com/apoydence/ledger/database"
	"github.com/apoydence/ledger/transaction"
)

func Combine(fs ...database.Filter) database.Filter {
	return FilterFunc(func(t *transaction.Transaction) bool {
		for _, f := range fs {
			if !f.Filter(t) {
				return false
			}
		}
		return true
	})
}
