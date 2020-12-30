package controller

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOPACKAGE/mock_$GOFILE

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/zhendong233/Books/internal/book/model"
	"github.com/zhendong233/Books/internal/book/service"
	"github.com/zhendong233/Books/pkg/httputil"
)

type BookController interface {
	GetBook(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
}

type bookController struct {
	bs service.BookService
}

func NewBookController(bs service.BookService) BookController {
	return &bookController{
		bs: bs,
	}
}

func (c *bookController) GetBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	spaceID := chi.URLParam(r, "bookId")
	req, err := c.bs.FindByID(ctx, spaceID)
	if err != nil {
		httputil.RespondError(ctx, w, err)
		return
	}
	httputil.RespondJSON(ctx, w, http.StatusOK, req)
}

func (c *bookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		httputil.RespondError(ctx, w, err)
		return
	}
	req, err := c.bs.CreateBook(ctx, &book)
	if err != nil {
		httputil.RespondError(ctx, w, err)
		return
	}
	httputil.RespondJSON(ctx, w, http.StatusOK, req)
}
