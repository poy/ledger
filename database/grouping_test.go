package database_test

import (
	"github.com/apoydence/ledger/database"
	"github.com/apoydence/ledger/transaction"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grouping", func() {

	var (
		mockFilter *mockFilter
		mockAgg    *mockAggregator

		start time.Time
		end   time.Time

		t1 *transaction.Transaction
		t2 *transaction.Transaction

		acc1 []*transaction.Account
		acc2 []*transaction.Account

		grouping *database.Grouping
	)

	BeforeEach(func() {
		start = transaction.NewDate(2015, 10, 10)
		end = transaction.NewDate(2015, 10, 12)

		acc1 = []*transaction.Account{
			{
				Name:  "some-acc-name-1",
				Value: 1,
			},
			{
				Name:  "some-acc-name-2",
				Value: 2,
			},
		}

		acc2 = []*transaction.Account{
			{
				Name:  "some-acc-name-2",
				Value: 3,
			},
		}

		t1 = &transaction.Transaction{
			Date: transaction.NewDate(2015, 10, 11),
			Title: &transaction.Title{
				Value: "some-title",
			},
			Accounts: &transaction.AccountList{
				Accounts: acc1,
			},
		}

		t2 = &transaction.Transaction{
			Date: transaction.NewDate(2015, 10, 11),
			Title: &transaction.Title{
				Value: "some-title",
			},
			Accounts: &transaction.AccountList{
				Accounts: acc2,
			},
		}

		mockFilter = newMockFilter()
		mockAgg = newMockAggregator()
		grouping = database.NewGrouping()
		grouping.Add(t1, t2)
	})

	It("groups based on account names", func() {
		mockAgg.resultCh <- "1234"
		mockAgg.resultCh <- "1234"
		results := grouping.Aggregate(start, end, nil, mockAgg)

		Expect(mockAgg.accountCh).To(HaveLen(2))
		var accs [][]*transaction.Account
		var acc []*transaction.Account
		Expect(mockAgg.accountCh).To(Receive(&acc))
		accs = append(accs, acc)
		Expect(mockAgg.accountCh).To(Receive(&acc))
		accs = append(accs, acc)

		Expect(accs).To(ContainElement([]*transaction.Account{
			acc1[0],
		}))

		Expect(accs).To(ContainElement([]*transaction.Account{
			acc1[1],
			acc2[0],
		}))

		Expect(results).To(HaveKeyWithValue("some-acc-name-1", []string{"1234"}))
		Expect(results).To(HaveKeyWithValue("some-acc-name-2", []string{"1234"}))
	})

})
