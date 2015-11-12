package transaction

import (
	"fmt"
)

type AccountList struct {
	Accounts []*Account
}

func (a *AccountList) Parse(line string) (string, error) {
	var err error
	for len(line) > 0 {
		account := new(Account)
		if line, err = account.Parse(line); err != nil {
			return "", err
		}
		a.Accounts = append(a.Accounts, account)
	}
	return "", a.reconcile()
}

func (a *AccountList) String() string {
	var result string

	for _, acc := range a.Accounts {
		result = fmt.Sprintf("%s\t%v\n", result, acc)
	}

	return result
}

func (a *AccountList) reconcile() error {
	var total float64
	for _, acc := range a.Accounts {
		total += acc.Value
	}

	if total != 0 {
		return fmt.Errorf("Does not reconcile. Off by $%-6.2f", total)
	}

	return nil
}
