package grpc

import (
	"errors"
	"github.com/oky-setiawan/stockbit-test/internal/entity"
	pb "github.com/oky-setiawan/stockbit-test/internal/handler/grpc/protos"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServer_GetMovieInfo(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, m testObject) {
			req := &pb.GetMovieInfoRequest{
				Keyword: "test",
				Page:    1,
			}

			data := &entity.SearchMovieResponse{Data: []entity.SearchMovieData{{}}}
			m.movieUC.On("SearchMovie", mock.Anything, mock.Anything).
				Return(data, nil)

			resp, err := m.handler.GetMovieInfo(m.ctx, req)
			g.Expect(err).Should(BeNil())
			g.Expect(resp.Data).Should(HaveLen(len(data.Data)))
		})
	})

	t.Run("Error SearchMovie", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, m testObject) {
			req := &pb.GetMovieInfoRequest{
				Keyword: "test",
				Page:    1,
			}

			data := &entity.SearchMovieResponse{Data: []entity.SearchMovieData{{}}}
			m.movieUC.On("SearchMovie", mock.Anything, mock.Anything).
				Return(data, errors.New("some error"))

			resp, err := m.handler.GetMovieInfo(m.ctx, req)
			g.Expect(err).Should(HaveOccurred())
			g.Expect(resp).Should(BeNil())
		})
	})

	t.Run("Error InvalidArgument", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, m testObject) {
			req := &pb.GetMovieInfoRequest{}

			resp, err := m.handler.GetMovieInfo(m.ctx, req)
			g.Expect(err).Should(HaveOccurred())
			g.Expect(resp).Should(BeNil())
		})
	})
}
