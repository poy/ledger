package transaction

import (
	"fmt"
	"math"
	"strings"
)

type Transaction struct {
	Date     *Date
	Title    *Title
	Accounts *AccountList
}

func (t *Transaction) Parse(line string) (string, error) {
	t.Date = new(Date)
	t.Title = new(Title)
	t.Accounts = new(AccountList)

	var err error
	if line, err = t.Date.Parse(line); err != nil {
		return "", err
	}

	if line, err = t.Title.Parse(line); err != nil {
		return "", err
	}

	if line, err = t.Accounts.Parse(line); err != nil {
		return "", err
	}

	return "", nil
}

func (t *Transaction) Format(width int) string {
	var accs string
	for _, a := range t.Accounts.Accounts {
		numOfSpaces := width - (len(a.Name) + 3 + t.numberOfDigits(a.Value))
		accs = fmt.Sprintf("%s  %s%s$%-0.2f\n", accs, a.Name, strings.Repeat(" ", numOfSpaces), a.Value)
	}
	return fmt.Sprintf("%v %v\n%s", t.Date, t.Title, accs)
}

func (t *Transaction) MinimumWidth() int {
	titleLen := 11 + len(t.Title.Value)
	accLen := t.longestAccount(t.Accounts.Accounts)

	if titleLen > accLen {
		return titleLen
	}

	return accLen
}

func (t *Transaction) longestAccount(accs []*Account) int {
	var longest int
	for _, a := range accs {
		length := len(a.Name) + t.numberOfDigits(a.Value)
		if length > longest {
			longest = length
		}
	}

	return longest + 4
}

func (t *Transaction) numberOfDigits(value float64) int {
	if value < 0 {
		return int(math.Log10(-1*value)) + 5
	}

	return int(math.Log10(value)) + 4
}
