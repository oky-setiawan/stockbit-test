package movie

import (
	"context"
	"github.com/oky-setiawan/stockbit-test/mocks/mock_internal/mock_domain"
	"github.com/oky-setiawan/stockbit-test/mocks/mock_internal/mock_repository"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNew(t *testing.T) {
	g := NewGomegaWithT(t)
	m := New(&Opts{})
	g.Expect(m).Should(BeAssignableToTypeOf(&movieUsecase{}))
}

type testObject struct {
	ctx        context.Context
	uc         *movieUsecase
	omdbDomain *mock_domain.OMDBDomain
}

func doTest(t *testing.T, fn func(g *GomegaWithT, m testObject)) {
	g := NewGomegaWithT(t)

	logRepo := &mock_repository.LogRepository{}
	logRepo.On("LogAction", mock.Anything, mock.Anything).Return(nil)

	omdbDomain := &mock_domain.OMDBDomain{}

	uc := &movieUsecase{
		logRepo:    logRepo,
		omdbDomain: omdbDomain,
	}

	m := testObject{
		ctx:        context.Background(),
		uc:         uc,
		omdbDomain: omdbDomain,
	}
	fn(g, m)
}
