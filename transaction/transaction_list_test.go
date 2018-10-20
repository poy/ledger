package transaction_test

import (
	"sort"

	"github.com/poy/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TransactionList", func() {
	Describe("Sort Interface", func() {
		var (
			t1   *transaction.Transaction
			t2   *transaction.Transaction
			t3   *transaction.Transaction
			list transaction.TransactionList
		)

		BeforeEach(func() {
			t1 = &transaction.Transaction{
				Date: transaction.NewDate(2012, 1, 1),
			}
			t2 = &transaction.Transaction{
				Date: transaction.NewDate(2013, 1, 1),
			}
			t3 = &transaction.Transaction{
				Date: transaction.NewDate(2011, 1, 1),
			}
			list = transaction.TransactionList([]*transaction.Transaction{
				t1,
				t2,
				t3,
			})
		})

		It("is sortable", func() {
			sort.Sort(list)
			Expect(list[0]).To(Equal(t3))
			Expect(list[1]).To(Equal(t1))
			Expect(list[2]).To(Equal(t2))
		})
	})
})
