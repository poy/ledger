package transaction_test

import (
	. "github.com/apoydence/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Money", func() {
	Describe("Parse()", func() {
		It("parses dollar and cents", func() {
			money, err := ParseMoney("12.34")

			Expect(err).ToNot(HaveOccurred())
			Expect(money).To(BeEquivalentTo(1234))
		})

		It("parses negative amounts", func() {
			money, err := ParseMoney("-12.34")

			Expect(err).ToNot(HaveOccurred())
			Expect(money).To(BeEquivalentTo(-1234))
		})

		It("returns an error for invalid dollar value", func() {
			_, err := ParseMoney("12x34")

			Expect(err).To(HaveOccurred())
		})

		It("returns an error for invalid cents value", func() {
			_, err := ParseMoney("12.x34")

			Expect(err).To(HaveOccurred())
		})
	})
	Describe("String()", func() {
		It("returns a human readable monetary value", func() {
			Expect(Money(12345).String()).To(Equal("$123.45"))
		})

		It("returns a human readable monetary value for negatives values", func() {
			Expect(Money(-12345).String()).To(Equal("$-123.45"))
		})
	})
})
