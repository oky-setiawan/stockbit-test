package core

import (
	"github.com/oky-setiawan/stockbit-test/internal/config"
	"github.com/oky-setiawan/stockbit-test/internal/usecase"
	"github.com/oky-setiawan/stockbit-test/lib/grace"
	"github.com/oky-setiawan/stockbit-test/lib/router"
)

func initRouter(dlv *Deliveries, cfg *config.MainConfig) *HTTPHandler {

	api := router.New(&router.Options{
		Prefix:  "/api",
		Timeout: cfg.Server.HTTPTimeout,
	})

	//movie
	api.GET("/movie", dlv.Get)

	return &HTTPHandler{
		cfg:     cfg,
		usecase: dlv.uc,
	}
}

type HTTPHandler struct {
	cfg         *config.MainConfig
	usecase     *usecase.Usecase
	listenErrCh chan error
}

// Run will run http server
func (c *HTTPHandler) Run() {
	c.listenErrCh <- grace.ServeHTTP(c.cfg.Server.HTTPPort, router.GetHandler())
}

// ListenErr will get err chan listen http
func (c *HTTPHandler) ListenErr() chan error {
	return c.listenErrCh
}
