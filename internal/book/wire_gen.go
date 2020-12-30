// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package book

import (
	"github.com/zhendong233/Books/internal/book/controller"
	"github.com/zhendong233/Books/internal/book/repository"
	"github.com/zhendong233/Books/internal/book/service"
	"github.com/zhendong233/Books/pkg/dbutil"
)

// Injectors from wire.go:

func WireBuild() (Master, error) {
	db, err := dbutil.NewDB()
	if err != nil {
		return nil, err
	}
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository, db)
	bookController := controller.NewBookController(bookService)
	bookBookRouter := NewRouter(bookController)
	bookMaster := NewMaster(db, bookBookRouter)
	return bookMaster, nil
}
