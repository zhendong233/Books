package testutil

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertResponseBody(t *testing.T, want interface{}, body io.Reader) {
	t.Helper()
	bw, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadAll(body)
	if err != nil {
		t.Fatal(err)
	}
	assert.JSONEq(t, string(bw), string(b))
}

func NewRequestAndRecorder(method, path string, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	return httptest.NewRecorder(), httptest.NewRequest(method, path, body)
}
