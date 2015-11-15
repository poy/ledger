package transaction

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	dateRegexpPattern = `^(\d{4})/(\d\d?)/(\d\d?)\s*([\w\W]*)$`
)

var (
	dateRegexp *regexp.Regexp
)

func init() {
	dateRegexp = regexp.MustCompile(dateRegexpPattern)
}

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) Parse(line string) (string, error) {
	parsed := dateRegexp.FindStringSubmatch(line)
	if len(parsed) == 0 {
		return "", fmt.Errorf("Didn't find date in line: %s", line)
	}

	d.Year = safelyParseInt(parsed[1])
	d.Month = safelyParseInt(parsed[2])
	d.Day = safelyParseInt(parsed[3])

	return parsed[4], nil
}

func (d *Date) GreaterThanEqualTo(other *Date) bool {
	if d.Year < other.Year {
		return false
	}

	if d.Month < other.Month {
		return false
	}

	if d.Day < other.Day {
		return false
	}

	return true
}

func (d *Date) String() string {
	return fmt.Sprintf("%04d/%02d/%02d", d.Year, d.Month, d.Day)
}

func safelyParseInt(value string) int {
	i, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("Invalid int(%s): %v", value, err))
	}
	return int(i)
}
