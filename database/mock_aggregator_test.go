package database_test

import "github.com/apoydence/ledger/transaction"

type mockAggregator struct {
	accountCh chan []*transaction.Account
	resultCh  chan string
}

func newMockAggregator() *mockAggregator {
	return &mockAggregator{
		accountCh: make(chan []*transaction.Account, 100),
		resultCh:  make(chan string, 100),
	}
}

func (m *mockAggregator) Aggregate(accounts []*transaction.Account) string {
	m.accountCh <- accounts
	return <-m.resultCh
}
