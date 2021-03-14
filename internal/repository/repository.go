package repository

import (
	"context"
	"github.com/oky-setiawan/stockbit-test/internal/entity"
)

type (
	LogRepository interface {
		LogAction(ctx context.Context, request *entity.LogActionRequest) (err error)
	}
)
