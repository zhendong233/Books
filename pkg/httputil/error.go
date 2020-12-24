package httputil

import (
	"net/http"

	"github.com/zhendong233/Books/pkg/bookserr"
)

var errStatusMap map[string]int

func init() {
	errStatusMapInit()
}

func errStatusMapInit() {
	errStatusMap = setErrorStatus(map[bookserr.Code]int{
		bookserr.Unexpected:           http.StatusInternalServerError,
		bookserr.Unauthorized:         http.StatusForbidden,
		bookserr.ResourceNotFound:     http.StatusNotFound,
		bookserr.InvalidRequestSchema: http.StatusBadRequest,
		bookserr.BadRequest:           http.StatusBadRequest,
		bookserr.Locked:               http.StatusLocked,
	})
}

func setErrorStatus(em map[bookserr.Code]int) map[string]int {
	m := make(map[string]int, len(em))
	for c, status := range em {
		m[c.String()] = status
	}
	return m
}

func ErrorToStatus(code bookserr.Code) int {
	v, ok := errStatusMap[code.String()]
	if !ok {
		return http.StatusInternalServerError
	}
	return v
}
