package omdb

import (
	"github.com/oky-setiawan/stockbit-test/internal/config"
	"net/http"
	"time"
)

type omdbDomain struct {
	HTTPClient *http.Client
	cfg        *config.OMDBConfig
}

func Init(cfg *config.PartnerConfig) *omdbDomain {
	return &omdbDomain{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		cfg:        &cfg.OMDB,
	}
}
