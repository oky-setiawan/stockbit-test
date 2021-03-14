package usecase

import (
	"context"
	"github.com/oky-setiawan/stockbit-test/internal/entity"
)

type (
	MovieUsecase interface {
		SearchMovie(ctx context.Context, request *entity.SearchMovieRequest) (response *entity.SearchMovieResponse, err error)
	}
)
