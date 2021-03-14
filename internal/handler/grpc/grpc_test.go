package grpc

import (
	"context"
	"github.com/oky-setiawan/stockbit-test/internal/usecase"
	"github.com/oky-setiawan/stockbit-test/mocks/mock_internal/mock_usecase"
	. "github.com/onsi/gomega"
	"testing"
)

type testObject struct {
	ctx     context.Context
	handler *Server
	movieUC *mock_usecase.MovieUsecase
}

func doTest(t *testing.T, fn func(g *GomegaWithT, m testObject)) {
	g := NewGomegaWithT(t)

	movieUC := &mock_usecase.MovieUsecase{}

	h := &Server{
		usecase: usecase.Usecase{
			MovieUsecase: movieUC,
		},
	}

	m := testObject{
		ctx:     context.Background(),
		handler: h,
		movieUC: movieUC,
	}
	fn(g, m)
}
