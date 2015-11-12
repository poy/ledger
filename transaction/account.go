package transaction

import (
	"fmt"
	"regexp"
	"strconv"
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
	Value float64
}

func (a *Account) Parse(line string) (string, error) {
	parsed := accountRegexp.FindStringSubmatch(line)
	if len(parsed) == 0 {
		return "", fmt.Errorf("Didn't find an account in the line: %s", line)
	}

	a.Name = parsed[1]
	a.Value = safelyParseFloat(parsed[2])
	return parsed[len(parsed)-1], nil
}

func (a *Account) String() string {
	return fmt.Sprintf("%s\t$%-6.2f", a.Name, a.Value)
}

func safelyParseFloat(value string) float64 {
	i, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(fmt.Sprintf("Invalid float64(%s): %v", value, err))
	}
	return float64(i)
}
