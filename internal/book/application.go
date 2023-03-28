package book

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/zhendong233/Books/pkg/logutil"
	booksmiddle "github.com/zhendong233/Books/pkg/middleware"
)

type Application interface {
	Run()
}

type application struct {
	db *sql.DB
	r  BookRouter
}

func NewApplication(db *sql.DB, r BookRouter) Application {
	return &application{
		r:  r,
		db: db,
	}
}

func (a *application) Run() {
	if err := a.run(); err != nil {
		logutil.Logger.Error().Err(err).Caller().Send()
		os.Exit(-1)
	}
}

func (a *application) run() error {
	defer func() {
		_ = a.db.Close()
	}()
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/books", booksmiddle.SetFakeSession(a.r.Handler()))
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
