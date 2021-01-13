package dbutil

import (
	"context"

	"github.com/johejo/sqlbuilder"
	"github.com/zhendong233/Books/pkg/ctxutil"
)

func BuidPaging(ctx context.Context, q string, args ...interface{}) (string, []interface{}) {
	var b sqlbuilder.Builder
	b.Append(q, args...)
	offset := ctxutil.OffSet(ctx)
	limit := ctxutil.Limit(ctx)
	if limit == 0 {
		limit = 10
	}
	b.Append(" LIMIT ?, ?", offset, limit)
	query, newargs := b.Build()
	return query, newargs
}
