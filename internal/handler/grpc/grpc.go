package grpc

import (
	"github.com/oky-setiawan/stockbit-test/internal/config"
	pb "github.com/oky-setiawan/stockbit-test/internal/handler/grpc/protos"
	"github.com/oky-setiawan/stockbit-test/internal/usecase"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp/reuseport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"os"
	"syscall"
)

//Server is type of grpc server
type Server struct {
	usecase usecase.Usecase
}

//Opts is option to create grpc server
type Opts struct {
	Cfg     *config.Config
	Usecase usecase.Usecase
}

type GRPCServer struct {
	s       *grpc.Server
	address string
}

//New will make new grpc server
func New(o *Opts) *GRPCServer {
	s := grpc.NewServer()
	pb.RegisterStockbitServer(s, &Server{
		usecase: o.Usecase,
	})
	reflection.Register(s)

	grpcServer := &GRPCServer{
		address: o.Cfg.Main.Server.GRPCPort,
		s:       s,
	}
	return grpcServer
}

// Start will start grpc server
func (s *GRPCServer) Start() error {
	l, err := reuseport.Listen("tcp4", s.address)
	if err != nil {
		return err
	}

	log.Println("starting grpc server on ", s.address)
	return s.s.Serve(l)
}

// CatchSignal will catch signal if grpc server need to stopped
func (s *GRPCServer) CatchSignal(sig os.Signal) {
	log.Println("grpc service got signal: ", sig)
	if sig == syscall.SIGHUP || sig == syscall.SIGTERM {
		s.s.GracefulStop()
	}
}
