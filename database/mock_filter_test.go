package database_test

import "github.com/apoydence/ledger/transaction"

type mockFilter struct {
	transactionCh chan *transaction.Transaction
	resultCh      chan []*transaction.Account
}

func newMockFilter() *mockFilter {
	return &mockFilter{
		transactionCh: make(chan *transaction.Transaction, 100),
		resultCh:      make(chan []*transaction.Account, 100),
	}
}

func (m *mockFilter) Filter(t *transaction.Transaction) []*transaction.Account {
	m.transactionCh <- t
	return <-m.resultCh
}
