package database_test

import (
	"github.com/apoydence/ledger/database"
	"github.com/apoydence/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Database", func() {

	var db *database.Database

	BeforeEach(func() {
		db = database.New()
	})

	Describe("Add() & Filter()", func() {
		var (
			ts         []*transaction.Transaction
			mockFilter *mockFilter
		)

		BeforeEach(func() {
			mockFilter = newMockFilter()
			ts = []*transaction.Transaction{
				{
					Accounts: &transaction.AccountList{
						Accounts: make([]*transaction.Account, 2),
					},
					Date: &transaction.Date{
						Year:  2015,
						Month: 10,
						Day:   3,
					},
				},
				{
					Accounts: &transaction.AccountList{
						Accounts: make([]*transaction.Account, 2),
					},
					Date: &transaction.Date{
						Year:  2014,
						Month: 10,
						Day:   3,
					},
				},
				{
					Accounts: &transaction.AccountList{
						Accounts: make([]*transaction.Account, 2),
					},
					Date: &transaction.Date{
						Year:  2014,
						Month: 11,
						Day:   3,
					},
				},
				{
					Accounts: &transaction.AccountList{
						Accounts: make([]*transaction.Account, 2),
					},
					Date: &transaction.Date{
						Year:  2015,
						Month: 1,
						Day:   3,
					},
				},
			}
		})

		It("filters the transactions by the time window", func() {
			mockFilter.resultCh <- make([]*transaction.Account, 1)
			mockFilter.resultCh <- make([]*transaction.Account, 1)
			db.Add(ts...)
			start := &transaction.Date{
				Year:  2015,
				Month: 1,
				Day:   1,
			}
			end := &transaction.Date{
				Year:  2015,
				Month: 12,
				Day:   31,
			}
			results := db.Query(start, end, nil)
			Expect(results).To(HaveLen(2))
			Expect(results).To(ContainElement(ts[0]))
			Expect(results).To(ContainElement(ts[3]))
		})

		It("filters the transactions by the time window and filter", func() {
			mockFilter.resultCh <- nil
			mockFilter.resultCh <- make([]*transaction.Account, 1)
			db.Add(ts...)
			start := &transaction.Date{
				Year:  2014,
				Month: 10,
				Day:   3,
			}
			end := &transaction.Date{
				Year:  2014,
				Month: 11,
				Day:   30,
			}
			results := db.Query(start, end, mockFilter)
			Expect(mockFilter.transactionCh).To(HaveLen(2))
			Expect(results).To(HaveLen(1))
			Expect(results).To(ContainElement(ts[2]))
		})
	})
})
