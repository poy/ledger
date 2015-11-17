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
	if cents < 0 {
		cents *= -1
	}
	return fmt.Sprintf("$%d.%d", dollars, cents)
}
