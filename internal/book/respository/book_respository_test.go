package respository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/zhendong233/Books/pkg/testutil"
)

type testRespository struct {
	db  *sql.DB
	ctx context.Context
	r   BookRespository
}

func newTestRespository(t *testing.T) *testRespository {
	if testing.Short() {
		t.Skip()
	}
	db := testutil.PrepareMySQL(t)
	ctx := context.Background()
	r := NewBookRespository(db)
	testutil.SetFakeTimeForMysql(t, db, testutil.TestTime)
	t.Cleanup(func() {
		_ = db.Close()
	})
	return &testRespository{
		db:  db,
		ctx: ctx,
		r:   r,
	}
}
