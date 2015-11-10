package database_test

import "github.com/apoydence/ledger/transaction"

type mockFilter struct {
	transactionCh chan *transaction.Transaction
	resultCh      chan bool
}

func newMockFilter() *mockFilter {
	return &mockFilter{
		transactionCh: make(chan *transaction.Transaction, 100),
		resultCh:      make(chan bool, 100),
	}
}

func (m *mockFilter) Filter(t *transaction.Transaction) bool {
	m.transactionCh <- t
	return <-m.resultCh
}
