package httputil

import (
	"net/http"
	"testing"

	"github.com/zhendong233/Books/pkg/bookserr"
)

func Test_ErrorToStatus(t *testing.T) {
	tests := map[bookserr.Code]int{
		bookserr.Unexpected:           http.StatusInternalServerError,
		bookserr.Unauthorized:         http.StatusForbidden,
		bookserr.ResourceNotFound:     http.StatusNotFound,
		bookserr.InvalidRequestSchema: http.StatusBadRequest,
		bookserr.BadRequest:           http.StatusBadRequest,
		bookserr.Locked:               http.StatusLocked,
	}
	for c, want := range tests {
		if actual := ErrorToStatus(c); actual != want {
			t.Errorf("failed to convert: actual=%v, expected=%v", actual, want)
		}
	}
}
