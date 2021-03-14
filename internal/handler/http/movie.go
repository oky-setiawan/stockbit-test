package http

import (
	"github.com/oky-setiawan/stockbit-test/internal/entity"
	"github.com/oky-setiawan/stockbit-test/internal/usecase"
	"github.com/oky-setiawan/stockbit-test/lib/response"
	"net/http"
	"strconv"
)

type movieDlv struct {
	movieUC usecase.MovieUsecase
}

func NewMovie(uc *usecase.Usecase) MovieDelivery {
	return &movieDlv{
		movieUC: uc.MovieUsecase,
	}
}

// Get will get movie info based on request
func (m *movieDlv) Get(r *http.Request) *response.JSONResponse {
	resp := response.NewJSONResponse()
	ctx := r.Context()
	urlQuery := r.URL.Query()
	page, err := strconv.Atoi(urlQuery.Get("page"))
	if err != nil {
		return resp.SetError(response.ErrBadRequest, "invalid page number")
	}

	data, err := m.movieUC.SearchMovie(ctx, &entity.SearchMovieRequest{
		Keyword: urlQuery.Get("keyword"),
		Page:    page,
	})
	if err != nil {
		return resp.SetError(response.ErrInternalServer)
	}

	return resp.SetData(data.Data)
}
