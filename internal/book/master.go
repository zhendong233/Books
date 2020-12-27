package book

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Master interface {
	Run()
}

type master struct {
	r BookRouter
}

func NewMaster(r BookRouter) Master {
	return &master{
		r: r,
	}
}

func (m *master) Run() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/books", m.r.Handler())
	})

	fmt.Println("Server listen at :8005")
	err := http.ListenAndServe(":8005", r)
	if err != nil {
		os.Exit(-1)
	}
}
