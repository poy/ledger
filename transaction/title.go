package transaction

import "strings"

type Title struct {
	Value string
}

func (t *Title) Parse(line string) (string, error) {
	values := strings.SplitN(line, "\n", 2)
	t.Value = values[0]
	if len(values) != 2 {
		return "", nil
	}

	return values[1], nil
}

func (t *Title) String() string {
	return t.Value
}
