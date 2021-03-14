package core

import (
	"github.com/oky-setiawan/stockbit-test/internal/handler/http"
	"github.com/oky-setiawan/stockbit-test/internal/usecase"
)

type Deliveries struct {
	uc *usecase.Usecase
	http.MovieDelivery
}

func initDelivery(domain *Domain, repo *Repositories) *Deliveries {

	uc := usecase.InitUsecase(&usecase.Opts{
		LogRepo:    repo.LogRepository,
		OMDBDomain: domain.OMDBDomain,
	})

	return &Deliveries{
		MovieDelivery: http.NewMovie(uc),
		uc:            uc,
	}
}
