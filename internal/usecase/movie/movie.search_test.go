package movie

import (
	"errors"
	"github.com/oky-setiawan/stockbit-test/internal/entity"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_movieUsecase_SearchMovie(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, m testObject) {
			req := &entity.SearchMovieRequest{}

			movieResp := &entity.GetMovieResponse{}
			m.omdbDomain.On("GetMovie", mock.Anything, mock.Anything).
				Return(movieResp, nil)

			resp, err := m.uc.SearchMovie(m.ctx, req)
			g.Expect(err).Should(BeNil())
			g.Expect(resp.Data).Should(HaveLen(len(movieResp.Data)))
		})
	})

	t.Run("Error GetMovie", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, m testObject) {
			req := &entity.SearchMovieRequest{}

			movieResp := &entity.GetMovieResponse{}
			m.omdbDomain.On("GetMovie", mock.Anything, mock.Anything).
				Return(movieResp, errors.New("some error"))

			resp, err := m.uc.SearchMovie(m.ctx, req)
			g.Expect(err).Should(HaveOccurred())
			g.Expect(resp).Should(BeNil())
		})
	})
}
