package transaction

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
	return "", nil
}
