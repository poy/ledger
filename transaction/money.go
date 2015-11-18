package transaction

import (
	"fmt"
	"strconv"
	"strings"
)

type Money int64

func ParseMoney(value string) (Money, error) {
	parts := strings.SplitN(value, ".", 2)

	dollars, err := strconv.ParseInt(parts[0], 10, 64)
	dollars *= 100
	if err != nil {
		return 0, err
	}

	if len(parts) == 1 {
		return Money(dollars), nil
	}

	cents, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return 0, err
	}

	if parts[0][0] == '-' {
		return Money(dollars - int64(cents)), nil
	}

	return Money(dollars + int64(cents)), nil
}

func (m Money) String() string {
	dollars := m / 100
	cents := m % 100

	if dollars < 0 {
		dollars *= -1
	}

	if cents < 0 {
		cents *= -1
	}

	var sign string
	if m < 0 {
		sign = "-"
	}

	return fmt.Sprintf("$%s%d.%02d", sign, dollars, cents)
}
