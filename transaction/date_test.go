package transaction_test

import (
	"github.com/apoydence/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Date", func() {
	Describe("Parse", func() {
		var date *transaction.Date

		BeforeEach(func() {
			date = new(transaction.Date)
		})

		Context("with nothing after the date", func() {
			It("returns the Date and an empty remaning", func() {
				line := "2015/10/12"
				remaining, err := date.Parse(line)

				Expect(err).ToNot(HaveOccurred())
				Expect(remaining).To(BeEmpty())
				Expect(date.Year).To(Equal(2015))
				Expect(date.Month).To(Equal(10))
				Expect(date.Day).To(Equal(12))
			})

			It("returns an error for an invalid date", func() {
				line := "2015/xx/12"
				_, err := date.Parse(line)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("with stuff after the date", func() {
			It("returns the Date and the remaniningg string", func() {
				line := "2015/10/12  remaining stuff"
				remaining, err := date.Parse(line)

				Expect(err).ToNot(HaveOccurred())
				Expect(remaining).To(Equal("remaining stuff"))
				Expect(date.Year).To(Equal(2015))
				Expect(date.Month).To(Equal(10))
				Expect(date.Day).To(Equal(12))
			})
		})
	})

})
