package filters_test

import (
	"github.com/poy/ledger/database"
	"github.com/poy/ledger/filters"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Store", func() {
	var mockFilterFactory filters.FilterFactoryFunc

	BeforeEach(func() {
		mockFilterFactory = filters.FilterFactoryFunc(func(string) (database.Filter, error) {
			return nil, nil
		})
	})

	It("has a map of filters", func() {
		filters.AddToStore("some-filter-name", mockFilterFactory)
		Expect(filters.Store()).To(HaveKey("some-filter-name"))
	})

	It("panics if the same name is added twice", func() {
		filters.AddToStore("some-other-name", nil)
		Expect(func() { filters.AddToStore("some-other-name", nil) }).To(Panic())
	})

	Describe("Fetch()", func() {
		It("returns all the matching filters", func() {
			filters.AddToStore("filter-1", mockFilterFactory)
			filters.AddToStore("filter-2", mockFilterFactory)
			filters, err := filters.Fetch("filter-1", "filter-2")

			Expect(err).ToNot(HaveOccurred())
			Expect(filters).To(HaveLen(2))
		})

		It("returns an error for an unknown name", func() {
			_, err := filters.Fetch("filter-3")
			Expect(err).To(HaveOccurred())
		})
	})

})
