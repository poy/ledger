package transaction_test

import (
	"github.com/poy/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
	Describe("Parse", func() {
		var account *transaction.Account

		BeforeEach(func() {
			account = new(transaction.Account)
		})

		Context("with additional lines", func() {
			It("parses the account line", func() {
				line := "Expenses:Auto:Gas     $10.12 \n\tLiabilities:MasterCard   $-10.12"
				remaining, err := account.Parse(line)

				Expect(err).ToNot(HaveOccurred())
				Expect(remaining).To(Equal("\tLiabilities:MasterCard   $-10.12"))
				Expect(account.Name).To(Equal("Expenses:Auto:Gas"))
				Expect(account.Value).To(BeEquivalentTo(1012))
			})
		})

		Context("without additional lines", func() {
			It("parses the account line", func() {
				line := "Expenses:Auto:Gas     $10.12"
				remaining, err := account.Parse(line)

				Expect(err).ToNot(HaveOccurred())
				Expect(remaining).To(BeZero())
				Expect(account.Name).To(Equal("Expenses:Auto:Gas"))
				Expect(account.Value).To(BeEquivalentTo(1012))
			})
		})
	})
})
