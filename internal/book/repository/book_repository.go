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
	Upsert(ctx context.Context, tx *sql.Tx, book *model.Book) error
	// FindListByID(ctx context.Context) ([]*model.Book, error)
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

func (r *bookRepository) Upsert(ctx context.Context, tx *sql.Tx, book *model.Book) error {
	const q = `INSERT INTO book (book_id, book_name, author, created_at) 
VALUES (?, ?, ?, NOW(3)) ON DUPLICATE KEY UPDATE book_name = ?, author = ?, created_at = NOW(3)`
	if _, err := tx.ExecContext(ctx, q, book.BookID, book.BookName,
		book.Author, book.BookName, book.Author); err != nil {
		return err
	}
	return nil
}

// func (r *bookRepository) FindListByID(ctx context.Context) ([]*model.Book, error) {
// 	const q = "SELECT book_id, book_name, author, created_at FROM book"
// }
