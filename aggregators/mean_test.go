package aggregators_test

import (
	"github.com/apoydence/ledger/aggregators"
	"github.com/apoydence/ledger/database"
	"github.com/apoydence/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mean", func() {

	var mean database.Aggregator

	BeforeEach(func() {
		mean = aggregators.NewMean()
	})

	It("returns the mean", func() {
		accs := []*transaction.Account{
			{
				Name:  "some-name-1",
				Value: 1000,
			},
			{
				Name:  "some-name-2",
				Value: 500,
			},
		}

		Expect(mean.Aggregate(accs)).To(Equal("$7.50"))
	})

	It("registers itself with the aggregator store", func() {
		Expect(aggregators.Store()).To(HaveKey("mean"))
	})

})
