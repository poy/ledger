package database_test

import "github.com/apoydence/ledger/transaction"

type mockAggregator struct {
	accountCh chan []*transaction.Account
	resultCh  chan int64
}

func newMockAggregator() *mockAggregator {
	return &mockAggregator{
		accountCh: make(chan []*transaction.Account, 100),
		resultCh:  make(chan int64, 100),
	}
}

func (m *mockAggregator) Aggregate(accounts []*transaction.Account) int64 {
	m.accountCh <- accounts
	return <-m.resultCh
}
