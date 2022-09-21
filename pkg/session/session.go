package session

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/zhendong233/Books/pkg/ctxutil"
)

func SetUserID(ctx context.Context, v string) context.Context {
	if v == "" {
		return ctx
	}
	return context.WithValue(ctx, ctxutil.CtxUserID, v)
}

func UserID(ctx context.Context) string {
	v, ok := ctx.Value(ctxutil.CtxUserID).(string)
	if !ok {
		log.Ctx(ctx).Error().Msg("no userID in context")
		return ""
	}
	return v
}
