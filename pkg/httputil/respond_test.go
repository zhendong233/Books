package httputil

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhendong233/Books/pkg/testutil"
)

func Test_RespondNoContent(t *testing.T) {
	rec := httptest.NewRecorder()
	RespondNoContent(context.Background(), rec)
	resp := rec.Result()
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

type testStruct struct {
	bookID string
}

func Test_RespondJSON(t *testing.T) {
	tests := []struct {
		name       string
		payload    interface{}
		wantStatus int
		wantErr    bool
	}{
		{
			name: "load",
			payload: testStruct{
				bookID: "bookid",
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			RespondJSON(context.Background(), rec, tt.wantStatus, tt.payload)
			resp := rec.Result()
			assert.Equal(t, tt.wantStatus, resp.StatusCode)
			if tt.wantErr {
				testutil.AssertResponseBody(t, tt.payload, resp.Body)
			}
		})
	}
}
