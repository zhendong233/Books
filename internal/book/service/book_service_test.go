package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zhendong233/Books/internal/book/mock/mock_repository"
	"github.com/zhendong233/Books/internal/book/model"
	"github.com/zhendong233/Books/pkg/testutil"
)

type testService struct {
	ctx context.Context
	db  *sql.DB
	bs  BookService
	br  *mock_repository.MockBookRepository
}

func newTestService(t *testing.T) *testService {
	ctrl := gomock.NewController(t)
	ctx := testutil.SetUpContextWithDefault()
	db := testutil.NewSQLMockDB(t)
	t.Cleanup(func() {
		_ = db.Close()
		ctrl.Finish()
	})

	br := mock_repository.NewMockBookRepository(ctrl)
	bs := NewBookService(br)
	return &testService{
		ctx: ctx,
		db:  db,
		bs:  bs,
		br:  br,
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
		expect  func(tbs *testService)
		want    *model.Book
		wantErr bool
	}{
		{
			name: "ok",
			expect: func(tbs *testService) {
				tbs.br.EXPECT().FindByID(gomock.Any(), testutil.TestBookID).Return(book, nil)
			},
			want:    book,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tbs := newTestService(t)
			tt.expect(tbs)
			got, err := tbs.bs.FindByID(tbs.ctx, testutil.TestBookID)
			if (err != nil) != tt.wantErr {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
