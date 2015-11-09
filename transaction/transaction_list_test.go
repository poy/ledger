package transaction_test

import (
	"github.com/apoydence/ledger/transaction"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TransactionList", func() {

	It("reads all the transactions", func() {
		reader := strings.NewReader(`
2015/10/12 Exxon
		Expenses:Auto:Gas         $10.00
		Liabilities:MasterCard   $-10.00


2015/10/12 Qdoba
		Expenses:Food:FastFood $21.50
		Liabilities:AmericanExpress $-21.50
`)
		transactionList, err := transaction.ReadList(reader)
		Expect(err).ToNot(HaveOccurred())
		Expect(transactionList).To(HaveLen(2))
		Expect(transactionList).To(ContainElement(
			&transaction.Transaction{
				Date: &transaction.Date{
					Year:  2015,
					Month: 10,
					Day:   12,
				},
				Title: &transaction.Title{
					Value: "Exxon",
				},
				Accounts: &transaction.AccountList{
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
				},
			}))

		Expect(transactionList).To(ContainElement(
			&transaction.Transaction{
				Date: &transaction.Date{
					Year:  2015,
					Month: 10,
					Day:   12,
				},
				Title: &transaction.Title{
					Value: "Qdoba",
				},
				Accounts: &transaction.AccountList{
					Accounts: []*transaction.Account{
						{
							Name:  "Expenses:Food:FastFood",
							Value: 21.50,
						},
						{
							Name:  "Liabilities:AmericanExpress",
							Value: -21.50,
						},
					},
				},
			}))

	})

})
