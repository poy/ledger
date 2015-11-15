package formatter_test

import (
	"strings"

	"github.com/apoydence/ledger/formatter"
	"github.com/apoydence/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Formatter", func() {

	var (
		f *formatter.Formatter
		t *transaction.Transaction
	)

	BeforeEach(func() {
		f = formatter.New()
	})

	var (
		title       string
		accountName string
	)

	JustBeforeEach(func() {
		t = &transaction.Transaction{
			Date: &transaction.Date{},
			Title: &transaction.Title{
				Value: title,
			},
			Accounts: &transaction.AccountList{
				Accounts: []*transaction.Account{
					{
						Name:  accountName,
						Value: -123.45,
					},
					{
						Name:  accountName,
						Value: 12.34,
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
				Expect(f.MinimumWidth(t)).To(Equal(len(title) + 11))
			})
		})

		Context("Long Account Name", func() {

			BeforeEach(func() {
				title = "short"
				accountName = "some-really-long-account-name"
			})

			It("gives the minimum width for a long title", func() {
				Expect(f.MinimumWidth(t)).To(Equal(len(accountName) + 9 + 2))
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
			result = f.Format(t, 40)
			lines = strings.Split(result, "\n")
			Expect(lines).To(HaveLen(4))
		})

		It("writes the date with a space and then the title", func() {
			Expect(lines[0]).To(MatchRegexp("^0000/00/00 some-title$"))
		})

		It("sizes the accounts with spaces to equal the given width", func() {
			Expect(lines[1]).To(MatchRegexp(`^  some-name {21}\$-123\.45$`))
			Expect(lines[2]).To(MatchRegexp(`^  some-name   {21}\$12\.34$`))
		})

	})
})
