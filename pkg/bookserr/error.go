package bookserr

import "fmt"

type Error interface {
	error
	Code() Code
}

type booksError struct {
	Origin  error  `json:"-"`
	ErrCode Code   `json:"code"`
	Msg     string `json:"msg"`
}

// 判断struct是否实现了interface没实现的话编译时会报错
var _ Error = (*booksError)(nil)

func New(err error, code Code, message string) *booksError {
	return &booksError{
		Origin:  err,
		ErrCode: code,
		Msg:     message,
	}
}

func (e *booksError) Error() string {
	if e.Origin == nil {
		return fmt.Sprintf("code=%s, msg=%s", e.ErrCode, e.Msg)
	}
	return fmt.Sprintf("originType=%T, origin=%s, code=%s, msg=%s", e.Origin, e.Origin.Error(), e.ErrCode, e.Msg)
}

func (e *booksError) Code() Code {
	return e.ErrCode
}
