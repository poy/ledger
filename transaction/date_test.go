package transaction_test

import (
	"github.com/poy/ledger/transaction"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Date", func() {
	Describe("DateToString()", func() {
		It("outputs a string that is the same as the parsed", func() {
			line := "2015/10/31"
			date, _, _ := transaction.ParseDate(line)
			Expect(transaction.DateToString(date)).To(Equal(line))
		})
	})

	Describe("Parse", func() {
		Context("with nothing after the date", func() {
			It("returns the Date and an empty remaning", func() {
				line := "2015/10/12"
				date, remaining, err := transaction.ParseDate(line)

				Expect(err).ToNot(HaveOccurred())
				Expect(remaining).To(BeEmpty())
				Expect(date.Year()).To(Equal(2015))
				Expect(date.Month()).To(BeEquivalentTo(10))
				Expect(date.Day()).To(Equal(12))
			})

			It("returns the Date with a single month digit", func() {
				line := "2015/1/12"
				date, remaining, err := transaction.ParseDate(line)

				Expect(err).ToNot(HaveOccurred())
				Expect(remaining).To(BeEmpty())
				Expect(date.Year()).To(Equal(2015))
				Expect(date.Month()).To(BeEquivalentTo(1))
				Expect(date.Day()).To(Equal(12))
			})

			It("returns the Date with a single dat digit", func() {
				line := "2015/10/1"
				date, remaining, err := transaction.ParseDate(line)

				Expect(err).ToNot(HaveOccurred())
				Expect(remaining).To(BeEmpty())
				Expect(date.Year()).To(Equal(2015))
				Expect(date.Month()).To(BeEquivalentTo(10))
				Expect(date.Day()).To(Equal(1))
			})

			It("returns an error for an invalid date", func() {
				line := "2015/xx/12"
				_, _, err := transaction.ParseDate(line)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("with stuff after the date", func() {
			It("returns the Date and the remaniningg string", func() {
				line := "2015/10/12  remaining stuff"
				date, remaining, err := transaction.ParseDate(line)

				Expect(err).ToNot(HaveOccurred())
				Expect(remaining).To(Equal("remaining stuff"))
				Expect(date.Year()).To(Equal(2015))
				Expect(date.Month()).To(BeEquivalentTo(10))
				Expect(date.Day()).To(Equal(12))
			})
		})
	})
})
