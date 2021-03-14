package core

import (
	"github.com/oky-setiawan/stockbit-test/internal/config"
	"github.com/oky-setiawan/stockbit-test/internal/domain"
	"github.com/oky-setiawan/stockbit-test/internal/domain/omdb"
)

type Domain struct {
	domain.OMDBDomain
}

func initDomain(cfg *config.PartnerConfig) *Domain {
	omdbDomain := omdb.Init(cfg)

	return &Domain{
		OMDBDomain: omdbDomain,
	}
}
