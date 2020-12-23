package bookserr

import "fmt"

type Error interface {
	error
	Code() Code
}

type booksError struct {
	origin error
	code   Code
	msg    string
}

var (
	// 判断struct是否实现了interface没实现的话编译时会报错
	_ Error = (*booksError)(nil)
)

func New(err error, code Code, message string) error {
	return &booksError{
		origin: err,
		code:   code,
		msg:    message,
	}
}

func (e *booksError) Error() string {
	if e.origin == nil {
		return fmt.Sprintf("code=%s, msg=%s", e.code, e.msg)
	}
	return fmt.Sprintf("originType=%T, origin=%s, code=%s, msg=%s", e.origin, e.origin.Error(), e.code, e.msg)
}

func (e *booksError) Code() Code {
	return e.code
}
