package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zhendong233/Books/internal/book/mock/mock_repository"
	"github.com/zhendong233/Books/internal/book/model"
	uuid "github.com/zhendong233/Books/pkg/mockableuuid"
	"github.com/zhendong233/Books/pkg/testutil"
)

type testService struct {
	ctx    context.Context
	db     *sql.DB
	bs     BookService
	br     *mock_repository.MockBookRepository
	mockDB sqlmock.Sqlmock
}

func newTestService(t *testing.T) *testService {
	ctrl := gomock.NewController(t)
	ctx := testutil.SetUpContextWithDefault()
	db, mockDB := testutil.NewSQLMock(t)
	t.Cleanup(func() {
		_ = db.Close()
		ctrl.Finish()
	})

	br := mock_repository.NewMockBookRepository(ctrl)
	bs := NewBookService(br, db)
	return &testService{
		ctx:    ctx,
		db:     db,
		bs:     bs,
		br:     br,
		mockDB: mockDB,
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

func Test_CreateBook(t *testing.T) {
	book := &model.Book{
		BookName:  "新书",
		Author:    "王某某",
		CreatedAt: testutil.TestTime,
	}
	bookResult := &model.Book{
		BookID:    "uuid",
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
				tbs.mockDB.ExpectBegin()
				tbs.br.EXPECT().Upsert(gomock.Any(), gomock.Any(), book).Return(nil)
				tbs.mockDB.ExpectCommit()
				tbs.br.EXPECT().FindByID(gomock.Any(), "uuid").Return(bookResult, nil)
			},
			want:    bookResult,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tbs := newTestService(t)
			defer uuid.Mock(t, "uuid")()
			tt.expect(tbs)
			got, err := tbs.bs.CreateBook(tbs.ctx, book)
			if (err != nil) != tt.wantErr {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
