package aggregators_test

import (
	"github.com/apoydence/ledger/aggregators"
	"github.com/apoydence/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Store", func() {
	var someAgg aggregators.AggregatorFunc

	BeforeEach(func() {
		someAgg = aggregators.AggregatorFunc(func([]*transaction.Account) string {
			return ""
		})
	})

	It("has a map of aggregators", func() {
		aggregators.AddToStore("some-agg-name", someAgg)
		Expect(aggregators.Store()).To(HaveKey("some-agg-name"))
	})

	It("panics if the same name is added twice", func() {
		aggregators.AddToStore("some-other-name", nil)
		Expect(func() { aggregators.AddToStore("some-other-name", nil) }).To(Panic())
	})

	Describe("Fetch()", func() {
		It("returns all the matching aggregators", func() {
			aggregators.AddToStore("agg-1", someAgg)
			aggregators.AddToStore("agg-2", someAgg)
			aggs, err := aggregators.Fetch("agg-1", "agg-2")

			Expect(err).ToNot(HaveOccurred())
			Expect(aggs).To(HaveLen(2))
		})

		It("returns an error for an unknown name", func() {
			_, err := aggregators.Fetch("agg-3")
			Expect(err).To(HaveOccurred())
		})
	})

})
