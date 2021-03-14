package domain

import (
	"context"
	"github.com/oky-setiawan/stockbit-test/internal/entity"
)

type (
	OMDBDomain interface {
		GetMovie(ctx context.Context, request *entity.GetMovieRequest) (response *entity.GetMovieResponse, err error)
	}
)
