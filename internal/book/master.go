package book

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/log"

	booksmiddle "github.com/zhendong233/Books/pkg/middleware"
)

type Master interface {
	Run()
}

type master struct {
	db *sql.DB
	r  BookRouter
}

func NewMaster(db *sql.DB, r BookRouter) Master {
	return &master{
		r:  r,
		db: db,
	}
}

func (m *master) Run() {
	if err := m.run(); err != nil {
		log.Print(err)
		os.Exit(-1)
	}
}

func (m *master) run() error {
	defer func() {
		_ = m.db.Close()
	}()
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/books", booksmiddle.SetFakeSession(m.r.Handler()))
	})

	fmt.Println("Server listen at :8005")
	server := &http.Server{
		Addr:              ":8005",
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           r,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
