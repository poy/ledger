package transaction_test

import (
	"github.com/apoydence/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transaction", func() {
	var t *transaction.Transaction

	BeforeEach(func() {
		t = new(transaction.Transaction)
	})

	It("parses a transaction", func() {
		line := "2015/10/12 Exxon\n\tExpenses:Auto:Gas         $10.00\n\tLiabilities:MasterCard   $-10.00"
		remaining, err := t.Parse(line)

		Expect(err).ToNot(HaveOccurred())
		Expect(remaining).To(BeZero())
		Expect(t.Date).To(Equal(&transaction.Date{
			Year:  2015,
			Month: 10,
			Day:   12,
		}))
		Expect(t.Title).To(Equal(&transaction.Title{
			Value: "Exxon",
		}))
		Expect(t.Accounts).To(Equal(&transaction.AccountList{
			Accounts: []*transaction.Account{
				{
					Name:  "Expenses:Auto:Gas",
					Value: 10,
				},
				{
					Name:  "Liabilities:MasterCard",
					Value: -10,
				},
			},
		}))
	})

})
