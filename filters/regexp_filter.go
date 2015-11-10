package filters

import (
	"github.com/apoydence/ledger/transaction"
	"regexp"
)

type Regexp struct {
	reg *regexp.Regexp
}

func NewRegexp(pattern string) (*Regexp, error) {
	reg, err := regexp.Compile(pattern)
	return &Regexp{
		reg: reg,
	}, err
}

func (r *Regexp) Filter(t *transaction.Transaction) bool {
	return r.reg.MatchString(t.Title.Value) || r.checkAccounts(t.Accounts)
}

func (r *Regexp) checkAccounts(accounts *transaction.AccountList) bool {
	if accounts == nil {
		return false
	}

	for _, acc := range accounts.Accounts {
		if r.reg.MatchString(acc.Name) {
			return true
		}
	}
	return false
}
