package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/zhendong233/Books/internal/book/service"
)

type BookController interface {
	GetBook(w http.ResponseWriter, r *http.Request)
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
	_, err := c.bs.FindByID(ctx, spaceID)
	if err != nil {

	}
}
