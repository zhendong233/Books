//go:generate go run github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package book

import (
	"github.com/google/wire"

	"github.com/zhendong233/Books/internal/book/controller"
	"github.com/zhendong233/Books/internal/book/repository"
	"github.com/zhendong233/Books/internal/book/service"
	"github.com/zhendong233/Books/pkg/dbutil"
)

func WireBuild() (Application, error) {
	wire.Build(
		dbutil.NewDB,
		repository.NewBookRepository,
		service.NewBookService,
		controller.NewBookController,
		NewRouter,
		NewApplication,
	)
	return nil, nil
}
