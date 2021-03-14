package core

import (
	"github.com/oky-setiawan/stockbit-test/internal/config"
	"github.com/oky-setiawan/stockbit-test/internal/handler/grpc"
)

type Core struct {
	dlv    *Deliveries
	driver *Driver
	domain *Domain
	cfg    *config.Config
}

type App struct {
	core *Core
	HTTP *HTTPHandler
	GRPC *grpc.GRPCServer
}

func newCore(cfg *config.Config) *Core {
	driver := initDriver(cfg)

	repo := initRepository(driver)

	domain := initDomain(&cfg.Partner)

	dlv := initDelivery(domain, repo)

	return &Core{
		dlv:    dlv,
		driver: driver,
		domain: domain,
		cfg:    cfg,
	}
}

// NewApp will create new App
func NewApp(cfg *config.Config) *App {
	core := newCore(cfg)

	http := initRouter(core.dlv, &core.cfg.Main)

	grpc := initGRPCServer(core.cfg, core.dlv)

	return &App{
		core: core,
		HTTP: http,
		GRPC: grpc,
	}
}
