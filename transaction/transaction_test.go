package transaction_test

import (
	"strings"

	"github.com/poy/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transaction", func() {
	var t *transaction.Transaction

	Describe("Parse", func() {
		BeforeEach(func() {
			t = new(transaction.Transaction)
		})

		It("parses a transaction", func() {
			line := "2015/10/12 Exxon\n\tExpenses:Auto:Gas         $10.00\n\tLiabilities:MasterCard   $-10.00"
			remaining, err := t.Parse(line)

			Expect(err).ToNot(HaveOccurred())
			Expect(remaining).To(BeZero())
			Expect(t.Date).To(Equal(transaction.NewDate(2015, 10, 12)))
			Expect(t.Title).To(Equal(&transaction.Title{
				Value: "Exxon",
			}))
			Expect(t.Accounts).To(Equal(&transaction.AccountList{
				Accounts: []*transaction.Account{
					{
						Name:  "Expenses:Auto:Gas",
						Value: 1000,
					},
					{
						Name:  "Liabilities:MasterCard",
						Value: -1000,
					},
				},
			}))
		})
	})

	Describe("Formatting", func() {
		var (
			title       string
			accountName string
		)

		JustBeforeEach(func() {
			t = &transaction.Transaction{
				Title: &transaction.Title{
					Value: title,
				},
				Accounts: &transaction.AccountList{
					Accounts: []*transaction.Account{
						{
							Name:  accountName,
							Value: -12345,
						},
						{
							Name:  accountName,
							Value: 1234,
						},
					},
				},
			}
		})

		Describe("MiniumWidth()", func() {

			Context("Long title", func() {

				BeforeEach(func() {
					title = "super-really-long-title"
					accountName = "short"
				})

				It("gives the minimum width for a long title", func() {
					Expect(t.MinimumWidth()).To(Equal(len(title) + 11))
				})
			})

			Context("Long Account Name", func() {

				BeforeEach(func() {
					title = "short"
					accountName = "some-really-long-account-name"
				})

				It("gives the minimum width for a long title", func() {
					// 2 for initial spaces
					// 11 for money value plus 2 spaces
					Expect(t.MinimumWidth()).To(Equal(len(accountName) + 10 + 2))
				})
			})
		})

		Describe("Format()", func() {
			var (
				result string
				lines  []string
			)

			BeforeEach(func() {
				title = "some-title"
				accountName = "some-name"
			})

			JustBeforeEach(func() {
				result = t.Format(40)
				lines = strings.Split(result, "\n")
				Expect(lines).To(HaveLen(4))
			})

			It("writes the date with a space and then the title", func() {
				Expect(lines[0]).To(MatchRegexp("^0001/01/01 some-title$"))
			})

			It("sizes the accounts with spaces to equal the given width", func() {
				Expect(lines[1]).To(MatchRegexp(`^  some-name {20}\$-123\.45$`))
				Expect(lines[2]).To(MatchRegexp(`^  some-name   {20}\$12\.34$`))
			})
		})
	})

})
