package filters_test

import (
	"github.com/apoydence/ledger/filters"
	"github.com/apoydence/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Combiner", func() {
	var (
		mockFilter1 *mockFilter
		mockFilter2 *mockFilter
		t           *transaction.Transaction
	)

	BeforeEach(func() {
		mockFilter1 = newMockFilter()
		mockFilter2 = newMockFilter()
		t = new(transaction.Transaction)
	})

	It("combines each filter in order", func() {
		mockFilter1.resultCh <- true
		mockFilter2.resultCh <- true
		result := filters.Combine(mockFilter1, mockFilter2)

		Expect(result).ToNot(BeNil())
		Expect(result.Filter(t)).To(BeTrue())
		Expect(mockFilter1.transactionCh).To(Receive(Equal(t)))
		Expect(mockFilter2.transactionCh).To(Receive(Equal(t)))
	})

	It("stops after a filter returns false", func() {
		mockFilter1.resultCh <- false
		mockFilter2.resultCh <- true
		result := filters.Combine(mockFilter1, mockFilter2)

		Expect(result).ToNot(BeNil())
		Expect(result.Filter(t)).To(BeFalse())
		Expect(mockFilter1.transactionCh).To(Receive(Equal(t)))
		Expect(mockFilter2.transactionCh).ToNot(Receive())
	})
})
