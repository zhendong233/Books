package book

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/zhendong233/Books/internal/book/controller"
)

type BookRouter interface {
	Handler() http.Handler
}

type bookRouter struct {
	controller controller.BookController
}

func NewRouter(controller controller.BookController) BookRouter {
	return &bookRouter{
		controller: controller,
	}
}

func (br *bookRouter) Handler() http.Handler {
	mux := chi.NewRouter()
	mux.Route("/book", func(r chi.Router) {
		r.Post("/", br.controller.CreateBook)
		r.Get("/{bookId}", br.controller.GetBook)
		r.Put("/{bookId}", br.controller.UpdateBook)
	})
	return mux
}
