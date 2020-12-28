package controller

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOPACKAGE/mock_$GOFILE

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/zhendong233/Books/internal/book/service"
	"github.com/zhendong233/Books/pkg/httputil"
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
	res, err := c.bs.FindByID(ctx, spaceID)
	if err != nil {
		log.Print(err)
		httputil.RespondError(ctx, w, err)
		return
	}
	httputil.RespondJSON(ctx, w, http.StatusOK, res)
}
