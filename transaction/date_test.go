package transaction_test

import (
	"github.com/apoydence/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Date", func() {

	var date *transaction.Date

	BeforeEach(func() {
		date = new(transaction.Date)
	})

	Describe("Parse", func() {
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

	Describe("GreaterThan()", func() {
		var (
			a *transaction.Date
			b *transaction.Date
		)

		BeforeEach(func() {
			a = &transaction.Date{
				Year:  2015,
				Month: 10,
				Day:   12,
			}
			b = &transaction.Date{
				Year:  2014,
				Month: 10,
				Day:   12,
			}
		})

		It("returns true if the given date is before", func() {
			Expect(a.GreaterThanEqualTo(b)).To(BeTrue())
		})

		It("returns false if the given date is after", func() {
			Expect(b.GreaterThanEqualTo(a)).To(BeFalse())
		})

		It("returns true if the given date is equal", func() {
			Expect(a.GreaterThanEqualTo(a)).To(BeTrue())
		})
	})

})
