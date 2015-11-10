package filters_test

import (
	"github.com/apoydence/ledger/filters"
	"github.com/apoydence/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegexpFilter", func() {
	var (
		t *transaction.Transaction
	)

	Context("Matching Title", func() {
		BeforeEach(func() {
			t = &transaction.Transaction{
				Title: &transaction.Title{
					Value: "xxyyzz",
				},
				Accounts: &transaction.AccountList{
					Accounts: []*transaction.Account{
						{
							Name: "Expenses:Food:FastFood",
						},
						{
							Name: "Liabilities:Visa",
						},
					},
				},
			}
		})

		It("returns true for a title that matches", func() {
			filter := filters.NewRegexp(`xxy{2}zz`)
			Expect(filter.Filter(t)).To(BeTrue())
		})

		It("returns false for a non-matching title", func() {
			filter := filters.NewRegexp(`something-else`)
			Expect(filter.Filter(t)).To(BeFalse())
		})
	})

	Context("Matching Accounts", func() {
		It("returns true for the first account name that matches", func() {
			filter := filters.NewRegexp(`Expenses\:\w+\:FastFood`)
			Expect(filter.Filter(t)).To(BeTrue())
		})

		It("returns true for the second account name that matches", func() {
			filter := filters.NewRegexp(`Visa`)
			Expect(filter.Filter(t)).To(BeTrue())
		})

		It("returns false for non-matching account names", func() {
			filter := filters.NewRegexp(`MasterCard`)
			Expect(filter.Filter(t)).To(BeFalse())
		})
	})
})
