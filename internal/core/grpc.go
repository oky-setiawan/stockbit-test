package core

import (
	"github.com/oky-setiawan/stockbit-test/internal/config"
	"github.com/oky-setiawan/stockbit-test/internal/handler/grpc"
	log "github.com/sirupsen/logrus"
)

func initGRPCServer(cfg *config.Config, dlv *Deliveries) *grpc.GRPCServer {
	grpcServer := grpc.New(&grpc.Opts{
		Cfg:     cfg,
		Usecase: *dlv.uc,
	})

	go func() {
		if err := grpcServer.Start(); err != nil {
			log.Fatalf("failed to start grpc server: %v", err.Error())
		}
	}()

	return grpcServer
}
