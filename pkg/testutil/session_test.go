package testutil

import (
	"testing"

	"github.com/zhendong233/Books/pkg/books"
	"github.com/zhendong233/Books/pkg/session"
)

func Test_SetUpContext(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := SetUpContext("user-id")
	if userID := session.UserID(ctx); userID != "user-id" {
		t.Fatalf("set up context fail got userID= %s", userID)
	}
	ctx = SetUpContextWithDefault()
	if userID := session.UserID(ctx); userID != books.DefaultAdmin {
		t.Fatalf("set up context with default fail got userID= %s", userID)
	}
}
