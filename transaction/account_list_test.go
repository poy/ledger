package transaction_test

import (
	"github.com/poy/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AccountList", func() {
	Describe("Parse", func() {
		var accountList *transaction.AccountList

		BeforeEach(func() {
			accountList = new(transaction.AccountList)
		})

		It("reads all the account lines", func() {
			line := "\tExpenses:Auto:Gas         $10.90\n\tLiabilities:MasterCard   $-10.90\n"
			remaining, err := accountList.Parse(line)
			Expect(err).ToNot(HaveOccurred())
			Expect(remaining).To(BeZero())
			Expect(accountList.Accounts).To(HaveLen(2))
			Expect(accountList.Accounts).To(ContainElement(&transaction.Account{
				Name:  "Expenses:Auto:Gas",
				Value: 1090,
			}))
			Expect(accountList.Accounts).To(ContainElement(&transaction.Account{
				Name:  "Liabilities:MasterCard",
				Value: -1090,
			}))

		})

		It("returns an error if the accounts don't reconcile", func() {
			line := "\tExpenses:Auto:Gas         $20.90\n\tLiabilities:MasterCard   $-10.90"
			_, err := accountList.Parse(line)
			Expect(err).To(HaveOccurred())
		})

	})
})
