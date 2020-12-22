package service

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOPACKAGE/mock_$GOFILE

import (
	"context"
	"errors"

	"github.com/zhendong233/Books/internal/book/model"
	"github.com/zhendong233/Books/internal/book/repository"
	"github.com/zhendong233/Books/pkg/books"
	"github.com/zhendong233/Books/pkg/session"
)

type BookService interface {
	FindByID(ctx context.Context, bookID string) (*model.Book, error)
}

type bookService struct {
	br repository.BookRepository
}

func NewBookService(br repository.BookRepository) BookService {
	return &bookService{
		br: br,
	}
}

func (s *bookService) FindByID(ctx context.Context, bookID string) (*model.Book, error) {
	userID := session.UserID(ctx)
	if userID != books.DefaultAdmin {
		return nil, errors.New("this user can not find book")
	}
	return s.br.FindByID(ctx, bookID)
}
