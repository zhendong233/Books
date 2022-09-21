package bookserr

import (
	"encoding/json"
	"errors"
	"fmt"
)

type code uint

const (
	Unexpected code = iota + 1
	Unauthorized
	ResourceNotFound
	InvalidRequestSchema
	BadRequest
	Locked
)

type Code interface {
	fmt.Stringer
	json.Marshaler
}

func (c code) String() string {
	return fmt.Sprintf("books-%02d", c)
}

func (c code) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.String())), nil
}

func As(err error, c Code) bool {
	var booksErr Error
	if errors.As(err, &booksErr) && booksErr.Code() == c {
		return true
	}
	return false
}
