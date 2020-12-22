package testutil

import (
	"context"

	"github.com/zhendong233/Books/pkg/books"
	"github.com/zhendong233/Books/pkg/session"
)

func SetUpContext(userID string) context.Context {
	ctx := context.Background()
	ctx = session.SetUserID(ctx, userID)
	return ctx
}

func SetUpContextWithDefault() context.Context {
	return SetUpContext(books.DefaultAdmin)
}
