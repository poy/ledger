package transaction_test

import (
	"github.com/apoydence/ledger/transaction"

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
			line := "\tExpenses:Auto:Gas         $10.90\n\tLiabilities:MasterCard   $-10.90"
			remaining, err := accountList.Parse(line)
			Expect(err).ToNot(HaveOccurred())
			Expect(remaining).To(BeZero())
			Expect(accountList.Accounts).To(HaveLen(2))
			Expect(accountList.Accounts).To(ContainElement(&transaction.Account{
				Name:  "Expenses:Auto:Gas",
				Value: 10.90,
			}))
			Expect(accountList.Accounts).To(ContainElement(&transaction.Account{
				Name:  "Liabilities:MasterCard",
				Value: -10.90,
			}))

		})

	})
})
