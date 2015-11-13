package filters

import (
	"github.com/apoydence/ledger/database"
	"github.com/apoydence/ledger/transaction"
)

func CombineFilters(filters ...database.Filter) database.Filter {
	return FilterFunc(func(t *transaction.Transaction) []*transaction.Account {
		clone := *t
		cloneAccountList := *clone.Accounts
		clone.Accounts = &cloneAccountList

		for _, f := range filters {
			clone.Accounts.Accounts = f.Filter(&clone)
		}
		return clone.Accounts.Accounts
	})
}
