package filters

import (
	"github.com/apoydence/ledger/database"
	"github.com/apoydence/ledger/transaction"
	"regexp"
)

func init() {
	AddToStore("regexp", FilterFactoryFunc(NewRegexp))
}

type Regexp struct {
	reg *regexp.Regexp
}

func NewRegexp(pattern string) (database.Filter, error) {
	reg, err := regexp.Compile(pattern)
	return &Regexp{
		reg: reg,
	}, err
}

func (r *Regexp) Filter(t *transaction.Transaction) []*transaction.Account {
	if r.reg.MatchString(t.Title.Value) {
		return t.Accounts.Accounts
	}

	return r.checkAccounts(t.Accounts)
}

func (r *Regexp) checkAccounts(accounts *transaction.AccountList) []*transaction.Account {
	if accounts == nil {
		return nil
	}

	var results []*transaction.Account
	for _, acc := range accounts.Accounts {
		if r.reg.MatchString(acc.Name) {
			results = append(results, acc)
		}
	}
	return results
}
