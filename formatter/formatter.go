package formatter

import (
	"fmt"
	"math"
	"strings"

	"github.com/apoydence/ledger/transaction"
)

type Formatter struct {
}

func New() *Formatter {
	return &Formatter{}
}

func (f *Formatter) MinimumWidth(t *transaction.Transaction) int {
	titleLen := 11 + len(t.Title.Value)
	accLen := f.longestAccount(t.Accounts.Accounts)

	if titleLen > accLen {
		return titleLen
	}

	return accLen
}

func (f *Formatter) Format(t *transaction.Transaction, width int) string {
	var accs string
	for _, a := range t.Accounts.Accounts {
		numOfSpaces := width - (len(a.Name) + 3 + f.numberOfDigits(a.Value))
		accs = fmt.Sprintf("%s  %s%s$%-0.2f\n", accs, a.Name, strings.Repeat(" ", numOfSpaces), a.Value)
	}
	return fmt.Sprintf("%v %v\n%s", t.Date, t.Title, accs)
}

func (f *Formatter) longestAccount(accs []*transaction.Account) int {
	var longest int
	for _, a := range accs {
		length := len(a.Name) + f.numberOfDigits(a.Value)
		if length > longest {
			longest = length
		}
	}

	return longest + 4
}

func (f *Formatter) numberOfDigits(value float64) int {
	if value < 0 {
		return int(math.Log10(-1*value)) + 5
	}

	return int(math.Log10(value)) + 4
}
