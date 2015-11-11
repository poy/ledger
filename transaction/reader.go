package transaction

import (
	"bufio"
	"fmt"
	"io"
)

type Reader struct {
	scanner     *bufio.Scanner
	currentLine int64
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{
		currentLine: -1,
		scanner:     bufio.NewScanner(reader),
	}
}

func (t *Reader) Next() (*Transaction, *Error) {
	var (
		block string
	)

	for t.scanner.Scan() {
		t.currentLine++
		if len(t.scanner.Text()) > 0 && (len(block) == 0 || isSpace(t.scanner.Text()[0])) {
			block = fmt.Sprintf("%s%s\n", block, t.scanner.Text())
			continue
		}

		tr, err := parseTransaction(block)
		if err != nil {
			return nil, NewError(err, t.currentLine)
		}

		if tr != nil {
			return tr, nil
		}

		block = ""
	}

	tr, err := parseTransaction(block)
	if err != nil {
		return nil, NewError(err, t.currentLine)
	}

	return tr, NewError(t.scanner.Err(), t.currentLine)
}

func isSpace(char byte) bool {
	return char == byte(' ') || char == byte('\t')
}

func parseTransaction(block string) (*Transaction, error) {
	if len(block) <= 0 {
		return nil, nil
	}

	t := new(Transaction)
	if _, err := t.Parse(block); err != nil {
		return nil, err
	}

	return t, nil
}
