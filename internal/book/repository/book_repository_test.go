package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zhendong233/Books/internal/book/model"
	"github.com/zhendong233/Books/pkg/testutil"
)

type testRepository struct {
	db  *sql.DB
	ctx context.Context
	r   BookRepository
}

func newTestRepository(t *testing.T) *testRepository {
	if testing.Short() {
		t.Skip()
	}
	db := testutil.PrepareMySQL(t)
	ctx := testutil.SetUpContextWithDefault()
	r := NewBookRepository(db)
	testutil.SetFakeTimeForMysql(t, db, testutil.TestTime)
	t.Cleanup(func() {
		_ = db.Close()
	})
	return &testRepository{
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
		CreatedAt: testutil.TestTime,
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
			tbr := newTestRepository(t)
			testutil.ExecSQLFile(t, tbr.db, "./testdata/test_repository.sql")
			got, err := tbr.r.FindByID(tbr.ctx, tt.bookID)
			if (err != nil) != tt.wantErr {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_Upsert(t *testing.T) {
	book := &model.Book{
		BookID:    testutil.TestBookID,
		BookName:  "改名了",
		Author:    "改名了",
		CreatedAt: testutil.TestTime,
	}
	newBook := &model.Book{
		BookID:    "book-id-1",
		BookName:  "最新书",
		Author:    "刘某某",
		CreatedAt: testutil.TestTime,
	}
	tests := []struct {
		name    string
		book    *model.Book
		want    *model.Book
		wantErr bool
	}{
		{
			name:    "ok: insert new book",
			book:    newBook,
			want:    newBook,
			wantErr: false,
		},
		{
			name:    "ok: update the book",
			book:    book,
			want:    book,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tbr := newTestRepository(t)
			testutil.ExecSQLFile(t, tbr.db, "./testdata/test_repository.sql")
			tx := testutil.BeginTx(t, tbr.db)
			if err := tbr.r.Upsert(tbr.ctx, tx, tt.book); err != nil {
				testutil.RollBackTx(t, tx)
				t.Fatal(err)
			}
			testutil.CommitTx(t, tx)
			got, err := tbr.r.FindByID(tbr.ctx, tt.book.BookID)
			if (err != nil) != tt.wantErr {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
