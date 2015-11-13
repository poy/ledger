package filters_test

import (
	"github.com/apoydence/ledger/database"
	"github.com/apoydence/ledger/filters"
	"github.com/apoydence/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CombineFilters", func() {
	var (
		someTransaction *transaction.Transaction
		filter          database.Filter
		mockFilter1     *mockFilter
		mockFilter2     *mockFilter
	)

	BeforeEach(func() {
		someTransaction = &transaction.Transaction{
			Accounts: &transaction.AccountList{},
		}
		mockFilter1 = newMockFilter()
		mockFilter2 = newMockFilter()
		filter = filters.CombineFilters(mockFilter1, mockFilter2)
	})

	It("uses the first filter then the second", func() {
		mockFilter1.resultCh <- []*transaction.Account{
			{
				Name: "some-name-1",
			},
		}
		expectedResults := []*transaction.Account{
			{
				Name: "some-name-2",
			},
		}

		mockFilter2.resultCh <- expectedResults
		before := *someTransaction
		acBefore := *someTransaction.Accounts
		results := filter.Filter(someTransaction)
		Expect(mockFilter1.transactionCh).To(HaveLen(1))
		Expect(mockFilter2.transactionCh).To(HaveLen(1))

		Expect(someTransaction).To(Equal(&before))
		Expect(someTransaction.Accounts).To(Equal(&acBefore))
		Expect(results).To(Equal(expectedResults))
	})
})
