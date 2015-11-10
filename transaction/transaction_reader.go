package transaction

import (
	"bufio"
	"fmt"
	"io"
)

type Reader struct {
	scanner *bufio.Scanner
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{
		scanner: bufio.NewScanner(reader),
	}
}

func (t *Reader) Next() (*Transaction, error) {
	var (
		block string
	)

	for t.scanner.Scan() {
		if len(t.scanner.Text()) > 0 && (len(block) == 0 || isSpace(t.scanner.Text()[0])) {
			block = fmt.Sprintf("%s%s\n", block, t.scanner.Text())
			continue
		}

		tr, err := parseTransaction(block)
		if err != nil {
			return nil, err
		}

		if tr != nil {
			return tr, nil
		}

		block = ""
	}

	tr, err := parseTransaction(block)
	if err != nil {
		return nil, err
	}

	return tr, t.scanner.Err()
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
