package ctxutil

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
)

func SetOffSet(ctx context.Context, offset int) context.Context {
	if offset == 0 {
		return ctx
	}
	return context.WithValue(ctx, CtxOffSet, offset)
}

func OffSet(ctx context.Context) int {
	offset, ok := ctx.Value(CtxOffSet).(int)
	if !(ok) {
		log.Error().Err(errors.New("no offset in this ctx"))
		return 0
	}
	return offset
}

func SetLimit(ctx context.Context, limit int) context.Context {
	if limit == 0 {
		return ctx
	}
	return context.WithValue(ctx, CtxLimit, limit)
}

func Limit(ctx context.Context) int {
	limit, ok := ctx.Value(CtxLimit).(int)
	if !ok {
		log.Error().Err(errors.New("no limit in this ctx"))
		return 0
	}
	return limit
}

func SetSort(ctx context.Context, sort string) context.Context {
	if sort == "" {
		return ctx
	}
	return context.WithValue(ctx, CtxSort, sort)
}

func Sort(ctx context.Context) string {
	sort, ok := ctx.Value(CtxSort).(string)
	if !ok {
		log.Error().Err(errors.New("no sort in this ctx"))
		return ""
	}
	return sort
}
