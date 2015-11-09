package transaction_test

import (
	"github.com/apoydence/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Title", func() {
	Describe("Parse", func() {

		var title *transaction.Title

		BeforeEach(func() {
			title = new(transaction.Title)
		})

		Context("has new line", func() {
			It("extracts the title until the newline", func() {
				line := "some title\nother stuff"
				remaining, err := title.Parse(line)
				Expect(err).ToNot(HaveOccurred())
				Expect(remaining).To(Equal("other stuff"))
				Expect(title.Value).To(Equal("some title"))
			})
		})

		Context("does not have new line", func() {
			It("extracts the title", func() {
				line := "some title"
				remaining, err := title.Parse(line)
				Expect(err).ToNot(HaveOccurred())
				Expect(remaining).To(BeZero())
				Expect(title.Value).To(Equal("some title"))
			})
		})
	})
})
