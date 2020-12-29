package middleware

import (
	"net/http"

	"github.com/zhendong233/Books/pkg/books"
	"github.com/zhendong233/Books/pkg/session"
)

func SetFakeSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = session.SetUserID(ctx, books.DefaultAdmin)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
