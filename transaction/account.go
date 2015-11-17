package transaction

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	accountRegexpPattern = `\s*([\w\:]+)\s+\$(\-?\d+(\.\d+)?)(\n+([\w\W]*))?`
)

var (
	accountRegexp *regexp.Regexp
)

func init() {
	accountRegexp = regexp.MustCompile(accountRegexpPattern)
}

type Account struct {
	Name  string
	Value Money
}

func (a *Account) Parse(line string) (string, error) {
	var err error

	lines := strings.SplitN(line, "\n", 2)
	parsed := accountRegexp.FindStringSubmatch(strings.Trim(lines[0], " \t"))
	if len(parsed) == 0 {
		return "", fmt.Errorf("Didn't find an account in the line: %s", line)
	}

	a.Name = parsed[1]
	a.Value, err = ParseMoney(parsed[2])
	if err != nil {
		return "", err
	}

	if len(lines) == 1 {
		return "", nil
	}
	return lines[1], nil
}

func (a *Account) String() string {
	return fmt.Sprintf("%s\t%v", a.Name, a.Value)
}
