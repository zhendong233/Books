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

func (r *bookRespository) FindByID(ctx context.Context, bookID string) (*model.Book, error) {
	const q = "SELECT book_id, book_name, author, created_at FROM book WHERE book_id = ?"
	var book model.Book
	if err := r.db.QueryRowxContext(ctx, q, bookID).StructScan(&book); err != nil {
		return nil, err
	}
	return &book, nil
}
