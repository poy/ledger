package filters

import (
	"regexp"

	"github.com/poy/ledger/database"
	"github.com/poy/ledger/transaction"
)

func init() {
	AddToStore("regexp-title", FilterFactoryFunc(func(pattern string) (database.Filter, error) {
		return NewRegexp(pattern, true)
	}))

	AddToStore("regexp-accounts", FilterFactoryFunc(func(pattern string) (database.Filter, error) {
		return NewRegexp(pattern, false)
	}))
}

type Regexp struct {
	reg        *regexp.Regexp
	matchTitle bool
}

func NewRegexp(pattern string, matchTitle bool) (database.Filter, error) {
	reg, err := regexp.Compile(pattern)
	return &Regexp{
		reg:        reg,
		matchTitle: matchTitle,
	}, err
}

func (r *Regexp) Filter(t *transaction.Transaction) []*transaction.Account {
	if r.matchTitle {
		return r.checkTitle(t)
	}

	return r.checkAccounts(t.Accounts)
}

func (r *Regexp) checkTitle(t *transaction.Transaction) []*transaction.Account {
	if r.reg.MatchString(t.Title.Value) {
		return t.Accounts.Accounts
	}

	return nil
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
