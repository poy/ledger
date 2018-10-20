package filters_test

import (
	"github.com/poy/ledger/filters"
	"github.com/poy/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegexpFilter", func() {
	var (
		t *transaction.Transaction
	)

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

	Context("Matching Title", func() {
		It("returns all the accounts for a title that matches", func() {
			filter, err := filters.NewRegexp(`xxy{2}zz`, true)
			Expect(err).ToNot(HaveOccurred())
			Expect(filter.Filter(t)).To(HaveLen(2))
		})

		It("returns false for a non-matching title", func() {
			filter, err := filters.NewRegexp(`Expenses`, true)
			Expect(err).ToNot(HaveOccurred())
			Expect(filter.Filter(t)).To(HaveLen(0))
		})
	})

	Context("Matching Accounts", func() {
		It("returns true for the first account name that matches", func() {
			filter, err := filters.NewRegexp(`Expenses\:\w+\:FastFood`, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(filter.Filter(t)).To(HaveLen(1))
		})

		It("returns true for the second account name that matches", func() {
			filter, err := filters.NewRegexp(`Visa`, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(filter.Filter(t)).To(HaveLen(1))
		})

		It("returns false for non-matching account names", func() {
			filter, err := filters.NewRegexp(`xxyyzz`, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(filter.Filter(t)).To(HaveLen(0))
		})
	})

	It("registers title version with the store", func() {
		Expect(filters.Fetch("regexp-title")).To(HaveLen(1))
	})

	It("registers accounts version with the store", func() {
		Expect(filters.Fetch("regexp-accounts")).To(HaveLen(1))
	})
})
