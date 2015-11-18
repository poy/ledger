package database

import (
	"github.com/apoydence/ledger/transaction"
)

func CombineFilters(filters ...Filter) Filter {
	return FilterFunc(func(t *transaction.Transaction) []*transaction.Account {
		clone := *t
		cloneAccountList := *clone.Accounts
		clone.Accounts = &cloneAccountList

		for _, f := range filters {
			if f == nil {
				continue
			}
			clone.Accounts.Accounts = f.Filter(&clone)
		}
		return clone.Accounts.Accounts
	})
}
