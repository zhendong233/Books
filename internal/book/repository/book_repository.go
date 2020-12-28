package repository

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOPACKAGE/mock_$GOFILE

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/zhendong233/Books/internal/book/model"
)

type BookRepository interface {
	FindByID(ctx context.Context, bookID string) (*model.Book, error)
}

type bookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *bookRepository) FindByID(ctx context.Context, bookID string) (*model.Book, error) {
	const q = "SELECT book_id, book_name, author, created_at FROM book WHERE book_id = ?"
	var book model.Book
	if err := r.db.QueryRowxContext(ctx, q, bookID).StructScan(&book); err != nil {
		return nil, err
	}
	return &book, nil
}
