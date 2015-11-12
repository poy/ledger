package transaction

import "fmt"

type Error struct {
	msg  string
	Line int64
}

func NewError(err error, line int64) *Error {
	return &Error{
		msg:  readErrMsg(err),
		Line: line,
	}
}

func (t *Error) Error() string {
	return fmt.Sprintf("%s (line %d)", t.msg, t.Line)
}

func readErrMsg(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
