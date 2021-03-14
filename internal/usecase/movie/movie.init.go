package movie

import (
	"github.com/oky-setiawan/stockbit-test/internal/domain"
	"github.com/oky-setiawan/stockbit-test/internal/repository"
)

type Opts struct {
	LogRepo    repository.LogRepository
	OMDBDomain domain.OMDBDomain
}

type movieUsecase struct {
	logRepo    repository.LogRepository
	omdbDomain domain.OMDBDomain
}

func New(o *Opts) *movieUsecase {
	return &movieUsecase{
		logRepo:    o.LogRepo,
		omdbDomain: o.OMDBDomain,
	}
}
