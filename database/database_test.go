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
		db = new(database.Database)
	})

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
				Date: transaction.NewDate(2015, 10, 3),
			},
			{
				Accounts: &transaction.AccountList{
					Accounts: make([]*transaction.Account, 2),
				},
				Date: transaction.NewDate(2014, 10, 3),
			},
			{
				Accounts: &transaction.AccountList{
					Accounts: make([]*transaction.Account, 2),
				},
				Date: transaction.NewDate(2014, 11, 3),
			},
			{
				Accounts: &transaction.AccountList{
					Accounts: make([]*transaction.Account, 2),
				},
				Date: transaction.NewDate(2015, 1, 3),
			},
		}
	})

	It("filters the transactions by the time window", func() {
		mockFilter.resultCh <- make([]*transaction.Account, 1)
		mockFilter.resultCh <- make([]*transaction.Account, 1)
		db.Add(ts...)
		start := transaction.NewDate(2015, 1, 1)
		end := transaction.NewDate(2015, 12, 31)
		results := db.Query(start, end, nil)

		Expect(results).To(HaveLen(2))
		Expect(results).To(ContainElement(ts[0]))
		Expect(results).To(ContainElement(ts[3]))
	})

	It("filters the transactions by the time window and filter", func() {
		mockFilter.resultCh <- nil
		mockFilter.resultCh <- make([]*transaction.Account, 1)
		db.Add(ts...)
		start := transaction.NewDate(2014, 10, 3)
		end := transaction.NewDate(2014, 11, 30)

		results := db.Query(start, end, mockFilter)
		Expect(mockFilter.transactionCh).To(HaveLen(2))
		Expect(results).To(HaveLen(1))
		Expect(results).To(ContainElement(ts[2]))
	})

	Context("With aggregation", func() {

		var (
			mockAggregator1 *mockAggregator
			mockAggregator2 *mockAggregator
		)

		BeforeEach(func() {
			mockAggregator1 = newMockAggregator()
			mockAggregator2 = newMockAggregator()
		})

		It("filters and aggregates the transactions", func() {
			mockFilter.resultCh <- nil
			mockFilter.resultCh <- make([]*transaction.Account, 1)
			mockAggregator1.resultCh <- "99"
			mockAggregator2.resultCh <- "101"
			db.Add(ts...)
			start := transaction.NewDate(2014, 10, 3)
			end := transaction.NewDate(2014, 11, 30)

			results, aggResults := db.Aggregate(start, end, mockFilter, mockAggregator1, mockAggregator2)

			Expect(mockFilter.transactionCh).To(HaveLen(2))
			Expect(results).To(HaveLen(1))
			Expect(results).To(ContainElement(ts[2]))
			Expect(aggResults).To(HaveLen(2))
			Expect(aggResults[0]).To(Equal("99"))
			Expect(aggResults[1]).To(Equal("101"))
		})
	})
})
