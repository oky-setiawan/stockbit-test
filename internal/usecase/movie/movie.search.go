package movie

import (
	"context"
	"github.com/oky-setiawan/stockbit-test/internal/constants"
	"github.com/oky-setiawan/stockbit-test/internal/entity"
	log "github.com/sirupsen/logrus"
)

// SearchMovie is usecase that will search movie info
func (u *movieUsecase) SearchMovie(ctx context.Context, request *entity.SearchMovieRequest) (response *entity.SearchMovieResponse, err error) {

	getMovieResp, err := u.omdbDomain.GetMovie(ctx, &entity.GetMovieRequest{
		Keyword: request.Keyword,
		Page:    request.Page,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err.Error(),
			"keyword": request.Keyword,
			"page":    request.Page,
		}).Errorln("[SearchMovie] failed SearchMovie")
		return nil, err
	}

	logRequest := &entity.LogActionRequest{
		Action:   "SearchMovie",
		Method:   constants.LogMethodAPIRequest,
		Request:  request,
		Response: response,
	}

	go u.logRepo.LogAction(context.Background(), logRequest)

	data := make([]entity.SearchMovieData, len(getMovieResp.Data))
	for i, v := range getMovieResp.Data {
		data[i] = entity.SearchMovieData{
			Title:  v.Title,
			Year:   v.Year,
			ImdbID: v.ImdbID,
			Type:   v.Type,
			Poster: v.Poster,
		}
	}

	response = &entity.SearchMovieResponse{
		Data: data,
	}

	return response, nil
}
