package http

import (
	"context"
	"errors"
	"github.com/oky-setiawan/stockbit-test/internal/entity"
	"github.com/oky-setiawan/stockbit-test/mocks/mock_internal/mock_usecase"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/url"
	"testing"
)

type testObject struct {
	ctx     context.Context
	dlv     *movieDlv
	movieUC *mock_usecase.MovieUsecase
}

func doTest(t *testing.T, fn func(g *GomegaWithT, m testObject)) {
	g := NewGomegaWithT(t)

	movieUC := &mock_usecase.MovieUsecase{}

	dlv := &movieDlv{
		movieUC: movieUC,
	}

	m := testObject{
		ctx:     context.Background(),
		dlv:     dlv,
		movieUC: movieUC,
	}
	fn(g, m)
}

func Test_movieDlv_Get(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, m testObject) {

			m.movieUC.On("SearchMovie", mock.Anything, mock.Anything).
				Return(&entity.SearchMovieResponse{}, nil)

			resp := m.dlv.Get(&http.Request{URL: &url.URL{RawQuery: "?keyword=test&page=1"}})
			g.Expect(resp.HTTPCode).Should(Equal(http.StatusOK))
		})
	})

	t.Run("500", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, m testObject) {

			m.movieUC.On("SearchMovie", mock.Anything, mock.Anything).
				Return(&entity.SearchMovieResponse{}, errors.New("some error"))

			resp := m.dlv.Get(&http.Request{URL: &url.URL{RawQuery: "?keyword=test&page=1"}})
			g.Expect(resp.HTTPCode).Should(Equal(http.StatusInternalServerError))
		})
	})

	t.Run("400", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, m testObject) {

			m.movieUC.On("SearchMovie", mock.Anything, mock.Anything).
				Return(&entity.SearchMovieResponse{}, nil)

			resp := m.dlv.Get(&http.Request{URL: &url.URL{RawQuery: "?keyword=test&page=test"}})
			g.Expect(resp.HTTPCode).Should(Equal(http.StatusBadRequest))
		})
	})
}
