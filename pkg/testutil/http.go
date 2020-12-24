package testutil

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertResponseBody(t *testing.T, want interface{}, body io.Reader) {
	t.Helper()
	bw, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	assert.JSONEq(t, string(bw), string(b))
}
