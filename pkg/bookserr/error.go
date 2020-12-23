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

func New(code Code, message string) error {
	return &booksError{
		code: code,
		msg:  message,
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
