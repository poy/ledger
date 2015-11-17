package aggregators_test

import (
	"github.com/apoydence/ledger/aggregators"
	"github.com/apoydence/ledger/database"
	"github.com/apoydence/ledger/transaction"

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

		Expect(count.Aggregate(accs)).To(BeEquivalentTo(2))
	})

	It("registers itself with the aggregator store", func() {
		Expect(aggregators.Store()).To(HaveKey("count"))
	})

})
