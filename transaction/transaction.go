package transaction

import "fmt"

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

func (t *Transaction) String() string {
	return fmt.Sprintf("%v %v\n%v", t.Date, t.Title, t.Accounts)
}
