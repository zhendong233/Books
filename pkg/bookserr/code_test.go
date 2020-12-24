package bookserr

import (
	"reflect"
	"testing"
)

func Test_code(t *testing.T) {
	var actul, want string
	want = "books-01"
	actul = Unexpected.String()
	if actul != want {
		t.Errorf("failed to convert code to string; want = %s actul = %s", want, actul)
	}
	b, err := Unexpected.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	bw := []byte(`"books-01"`)
	if !reflect.DeepEqual(b, bw) {
		t.Errorf("failed to marshal json: want = %s actul = %s", want, actul)
	}
}
