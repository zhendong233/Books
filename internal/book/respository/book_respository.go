package respository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/zhendong233/Books/internal/book/model"
)

type BookRespository interface {
	FindByID(ctx context.Context, bookID string) (*model.Book, error)
}

type bookRespository struct {
	db *sqlx.DB
}

func NewBookRespository(db *sql.DB) *bookRespository {
	return &bookRespository{
		db: sqlx.NewDb(db, "mysql"),
	}
}
