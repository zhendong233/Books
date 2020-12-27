//go:generate wire
//+build wireinject

package book

import (
	"github.com/google/wire"

	"github.com/zhendong233/Books/internal/book/controller"
	"github.com/zhendong233/Books/internal/book/repository"
	"github.com/zhendong233/Books/internal/book/service"
	"github.com/zhendong233/Books/pkg/dbutil"
)

func WireBuild() (Master, error) {
	wire.Build(
		dbutil.NewDB,
		repository.NewBookRepository,
		service.NewBookService,
		controller.NewBookController,
		NewRouter,
		NewMaster,
	)
	return nil, nil
}
