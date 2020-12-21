package session

import (
	"context"

	"github.com/rs/zerolog/log"
)

type ctxUserIDType int

var (
	ctxUserID ctxUserIDType = 1
)

func SetUserID(ctx context.Context, v string) context.Context {
	if v == "" {
		return ctx
	}
	return context.WithValue(ctx, ctxUserID, v)
}

func UserID(ctx context.Context) string {
	v, ok := ctx.Value(ctxUserID).(string)
	if !ok {
		log.Ctx(ctx).Error().Msg("no userID in context")
		return ""
	}
	return v
}
