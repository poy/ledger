package filters

import "github.com/poy/ledger/database"

type FilterFactoryFunc func(arg string) (database.Filter, error)

func (f FilterFactoryFunc) Generate(arg string) (database.Filter, error) {
	return f(arg)
}
