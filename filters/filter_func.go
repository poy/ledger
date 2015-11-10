package filters

import "github.com/apoydence/ledger/transaction"

type FilterFunc func(*transaction.Transaction) bool

func (f FilterFunc) Filter(t *transaction.Transaction) bool {
	return f(t)
}
