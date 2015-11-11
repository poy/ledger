package database_test

import "github.com/apoydence/ledger/transaction"

type mockAggregator struct {
	accountCh chan []*transaction.Account
	resultCh  chan float64
}

func newMockAggregator() *mockAggregator {
	return &mockAggregator{
		accountCh: make(chan []*transaction.Account, 100),
		resultCh:  make(chan float64, 100),
	}
}

func (m *mockAggregator) Aggregate(accounts []*transaction.Account) float64 {
	m.accountCh <- accounts
	return <-m.resultCh
}
