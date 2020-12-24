package bookserr

import (
	"fmt"
	"testing"
)

func Test_error(t *testing.T) {
	newerr := New(nil, Unexpected, "new error")
	want := fmt.Sprintf("code=%s, msg=%s", Unexpected, "new error")
	if newerr.Error() != want {
		t.Errorf("failed creat newerr; want = %s actul = %s", want, newerr.Error())
	}
	if !As(newerr, Unexpected) {
		t.Error("failed creat newerr")
	}
}
