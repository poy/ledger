package transaction

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
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

func NewDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func ParseDate(line string) (time.Time, string, error) {
	parsed := dateRegexp.FindStringSubmatch(line)
	if len(parsed) == 0 {
		return time.Time{}, "", fmt.Errorf("Didn't find date in line: %s", line)
	}

	year := safelyParseInt(parsed[1])
	month := safelyParseInt(parsed[2])
	day := safelyParseInt(parsed[3])

	return NewDate(year, month, day), parsed[4], nil
}

func DateToString(d time.Time) string {
	return fmt.Sprintf("%04d/%02d/%02d", d.Year(), d.Month(), d.Day())
}

func safelyParseInt(value string) int {
	i, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("Invalid int(%s): %v", value, err))
	}
	return int(i)
}
