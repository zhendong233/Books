package service

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOPACKAGE/mock_$GOFILE

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/zhendong233/Books/internal/book/model"
	"github.com/zhendong233/Books/internal/book/repository"
	"github.com/zhendong233/Books/pkg/books"
	"github.com/zhendong233/Books/pkg/bookserr"
	uuid "github.com/zhendong233/Books/pkg/mockableuuid"
	"github.com/zhendong233/Books/pkg/session"
)

type BookService interface {
	FindByID(ctx context.Context, bookID string) (*model.Book, error)
	CreateBook(ctx context.Context, book *model.Book) (*model.Book, error)
}

type bookService struct {
	db *sqlx.DB
	br repository.BookRepository
}

func NewBookService(br repository.BookRepository, db *sql.DB) BookService {
	return &bookService{
		br: br,
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (s *bookService) FindByID(ctx context.Context, bookID string) (*model.Book, error) {
	userID := session.UserID(ctx)
	if userID != books.DefaultAdmin {
		return nil, bookserr.New(nil, bookserr.Unauthorized, "user can not find")
	}
	return s.br.FindByID(ctx, bookID)
}

func (s *bookService) CreateBook(ctx context.Context, book *model.Book) (*model.Book, error) {
	userID := session.UserID(ctx)
	if userID != books.DefaultAdmin {
		return nil, bookserr.New(nil, bookserr.Unauthorized, "user can not create book")
	}
	bookID := uuid.New().String()
	book.BookID = bookID
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return nil, err
	}
	if err := s.br.Upsert(ctx, tx, book); err != nil {
		return nil, tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return s.br.FindByID(ctx, bookID)
}
