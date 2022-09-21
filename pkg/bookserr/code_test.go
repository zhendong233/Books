package bookserr

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

func Test_As(t *testing.T) {
	tests := []struct {
		name string
		err  error
		code Code
		want bool
	}{
		{
			name: "ok true",
			err:  New(nil, Unexpected, ""),
			code: Unexpected,
			want: true,
		},
		{
			name: "ok: code not equal -- false",
			err:  New(nil, Unexpected, ""),
			code: Unauthorized,
			want: false,
		},
		{
			name: "ok: not books err -- false",
			err:  errors.New(""),
			code: Unauthorized,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := As(tt.err, tt.code)
			assert.Equal(t, tt.want, got)
		})
	}
}
