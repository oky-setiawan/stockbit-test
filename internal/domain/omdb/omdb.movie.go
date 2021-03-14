package omdb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/oky-setiawan/stockbit-test/internal/entity"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// getResponse will get entity.GetMovieResponse from omdbGetMovieResponse
func (o *omdbGetMovieResponse) getResponse() (response *entity.GetMovieResponse, err error) {
	if o.Response != responseOk {
		return nil, errors.New(o.Error)
	}

	response = &entity.GetMovieResponse{
		Data: []entity.GetMovieData{},
	}

	for _, v := range o.Data {
		data := entity.GetMovieData{
			Title:  v.Title,
			Year:   v.Year,
			ImdbID: v.ImdbID,
			Type:   v.Type,
			Poster: v.Poster,
		}
		response.Data = append(response.Data, data)
	}

	return response, nil
}

// GetMovie will get movie info from omdbDomain
func (o *omdbDomain) GetMovie(ctx context.Context, request *entity.GetMovieRequest) (response *entity.GetMovieResponse, err error) {

	httpReq, err := http.NewRequest(http.MethodGet, o.cfg.Host+o.cfg.GetMovieUrl, nil)
	if err != nil {
		return nil, err
	}

	//create url query
	urlQuery := &url.Values{}
	urlQuery.Add("apikey", o.cfg.AccessKey)
	urlQuery.Add("s", request.Keyword)
	urlQuery.Add("page", strconv.Itoa(request.Page))
	httpReq.URL.RawQuery = urlQuery.Encode()
	httpReq.WithContext(ctx)

	resp, err := o.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[omdbDomain][GetMovie] non 200 response")
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("[omdbDomain][GetMovie] failed ioutil.ReadAll, err: %v", err.Error())
	}

	httpResp := &omdbGetMovieResponse{}
	err = json.Unmarshal(b, httpResp)
	if err != nil {
		return nil, fmt.Errorf("[omdbDomain][GetMovie] failed  json.Unmarshal, err: %v", err.Error())
	}

	response, err = httpResp.getResponse()
	return response, err
}
