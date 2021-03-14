package usecase

import (
	"github.com/oky-setiawan/stockbit-test/internal/domain"
	"github.com/oky-setiawan/stockbit-test/internal/repository"
	"github.com/oky-setiawan/stockbit-test/internal/usecase/movie"
)

type Usecase struct {
	MovieUsecase
}

type Opts struct {
	LogRepo    repository.LogRepository
	OMDBDomain domain.OMDBDomain
}

func InitUsecase(o *Opts) *Usecase {
	movieUC := movie.New(&movie.Opts{
		LogRepo:    o.LogRepo,
		OMDBDomain: o.OMDBDomain,
	})
	return &Usecase{
		MovieUsecase: movieUC,
	}
}
