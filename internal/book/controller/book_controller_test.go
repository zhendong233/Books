package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zhendong233/Books/internal/book/mock/mock_service"
	"github.com/zhendong233/Books/internal/book/model"
	"github.com/zhendong233/Books/pkg/testutil"
)

type testController struct {
	bc  BookController
	bs  *mock_service.MockBookService
	ctx context.Context
	mux *chi.Mux
}

func newTestController(t *testing.T) *testController {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	bs := mock_service.NewMockBookService(ctrl)
	bc := NewBookController(bs)
	mux := chi.NewMux()
	return &testController{
		bc:  bc,
		bs:  bs,
		ctx: testutil.SetUpContextWithDefault(),
		mux: mux,
	}
}

func Test_GetBook(t *testing.T) {
	book := &model.Book{
		BookID:    testutil.TestBookID,
		BookName:  "新书",
		Author:    "王某某",
		CreatedAt: testutil.TestTime,
	}
	tests := []struct {
		name       string
		expect     func(tbc *testController)
		wantStatus int
		wantErr    bool
		wantBody   *model.Book
	}{
		{
			name: "ok",
			expect: func(tbc *testController) {
				tbc.bs.EXPECT().FindByID(gomock.Any(), testutil.TestBookID).Return(book, nil)
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
			wantBody:   book,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tbc := newTestController(t)
			tt.expect(tbc)
			tbc.mux.Get("/book/{bookId}", tbc.bc.GetBook)
			rec, req := testutil.NewRequestAndRecorder("GET", "/book/"+testutil.TestBookID, nil)
			tbc.mux.ServeHTTP(rec, req)
			assert.Equal(t, tt.wantStatus, rec.Code)
			if !tt.wantErr {
				testutil.AssertResponseBody(t, tt.wantBody, rec.Body)
			}
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
		BookID:    "book-id-1",
		BookName:  "新书",
		Author:    "王某某",
		CreatedAt: testutil.TestTime,
	}
	tests := []struct {
		name       string
		expect     func(tbc *testController)
		wantStatus int
		wantErr    bool
		wantBody   *model.Book
	}{
		{
			name: "ok",
			expect: func(tbc *testController) {
				tbc.bs.EXPECT().CreateBook(gomock.Any(), book).Return(bookResult, nil)
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
			wantBody:   bookResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tbc := newTestController(t)
			tt.expect(tbc)
			buf := new(bytes.Buffer)
			if err := json.NewEncoder(buf).Encode(book); err != nil {
				t.Fatal(err)
			}
			tbc.mux.Post("/book", tbc.bc.CreateBook)
			rec, req := testutil.NewRequestAndRecorder("POST", "/book", buf)
			tbc.mux.ServeHTTP(rec, req)
			assert.Equal(t, tt.wantStatus, rec.Code)
			if !tt.wantErr {
				testutil.AssertResponseBody(t, tt.wantBody, rec.Body)
			}
		})
	}
}
