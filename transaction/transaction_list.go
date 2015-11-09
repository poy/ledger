package transaction

import (
	"bufio"
	"fmt"
	"io"
)

type TransactionList []*Transaction

func ReadList(reader io.Reader) (TransactionList, error) {
	var (
		transactions TransactionList
		block        string
		err          error
	)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if len(scanner.Text()) > 0 && (len(block) == 0 || isSpace(scanner.Text()[0])) {
			block = fmt.Sprintf("%s%s\n", block, scanner.Text())
			continue
		}

		transactions, err = parseTransaction(block, transactions)
		if err != nil {
			return nil, err
		}

		block = ""
	}

	transactions, err = parseTransaction(block, transactions)
	if err != nil {
		return nil, err
	}

	return transactions, scanner.Err()
}

func isSpace(char byte) bool {
	return char == byte(' ') || char == byte('\t')
}

func parseTransaction(block string, tl TransactionList) (TransactionList, error) {
	if len(block) <= 0 {
		return tl, nil
	}

	t := new(Transaction)
	if _, err := t.Parse(block); err != nil {
		return nil, err
	}

	return append(tl, t), nil
}
