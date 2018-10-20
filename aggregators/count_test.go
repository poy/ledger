package aggregators_test

import (
	"github.com/poy/ledger/aggregators"
	"github.com/poy/ledger/database"
	"github.com/poy/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Count", func() {
	var count database.Aggregator

	BeforeEach(func() {
		count = aggregators.NewCount()
	})

	It("Sums all the account values", func() {
		accs := []*transaction.Account{
			{
				Name:  "some-name-1",
				Value: 1234,
			},
			{
				Name:  "some-name-2",
				Value: 5678,
			},
		}

		Expect(count.Aggregate(accs)).To(Equal("2"))
	})

	It("registers itself with the aggregator store", func() {
		Expect(aggregators.Store()).To(HaveKey("count"))
	})

})
