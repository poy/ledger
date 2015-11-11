package aggregators_test

import (
	"github.com/apoydence/ledger/aggregators"
	"github.com/apoydence/ledger/database"
	"github.com/apoydence/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sum", func() {

	var sum database.Aggregator

	BeforeEach(func() {
		sum = aggregators.NewSum()
	})

	It("Sums all the account values", func() {
		accs := []*transaction.Account{
			{
				Name:  "some-name-1",
				Value: 12.34,
			},
			{
				Name:  "some-name-2",
				Value: 56.78,
			},
		}

		Expect(sum.Aggregate(accs)).To(Equal(69.12))
	})

})
