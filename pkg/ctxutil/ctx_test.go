package ctxutil

import (
	"context"
	"testing"
)

func Test_SetOffSet(t *testing.T) {
	ctx := context.Background()
	ctx = SetOffSet(ctx, 1)
	offSet := OffSet(ctx)
	if offSet != 1 {
		t.Fatal("SetOffSet and OffSet fail")
	}
}

func Test_Limit(t *testing.T) {
	ctx := context.Background()
	ctx = SetLimit(ctx, 5)
	limit := Limit(ctx)
	if limit != 5 {
		t.Fatal("SetLimit and Limit fail")
	}
}

func Test_Sort(t *testing.T) {
	ctx := context.Background()
	ctx = SetSort(ctx, "desc")
	sort := Sort(ctx)
	if sort != "desc" {
		t.Fatal("SetLimit and Limit fail")
	}
}
