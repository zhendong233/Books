package respository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zhendong233/Books/internal/book/model"
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

func Test_FindByID(t *testing.T) {
	book := &model.Book{
		BookID:    testutil.TestBookID,
		BookName:  "新书",
		Author:    "王某某",
		CreatedAt: testutil.ATestTime,
	}
	tests := []struct {
		name    string
		bookID  string
		want    *model.Book
		wantErr bool
	}{
		{
			name:    "ok",
			bookID:  testutil.TestBookID,
			want:    book,
			wantErr: false,
		},
		{
			name:    "ok no this book",
			bookID:  "123456",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tbr := newTestRespository(t)
			testutil.ExecSQLFile(t, tbr.db, "./testdata/test_repository.sql")
			got, err := tbr.r.FindByID(tbr.ctx, tt.bookID)
			if (err != nil) != tt.wantErr {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
