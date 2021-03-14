package grpc

import (
	"context"
	"github.com/oky-setiawan/stockbit-test/internal/entity"
	pb "github.com/oky-setiawan/stockbit-test/internal/handler/grpc/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetMovieInfo will get info movie based on request params
func (s Server) GetMovieInfo(ctx context.Context, request *pb.GetMovieInfoRequest) (*pb.GetMovieInfoResponse, error) {

	if request == nil || request.GetKeyword() == "" || request.GetPage() < 1 {
		return nil, status.Error(codes.InvalidArgument, "invalid request, please check your request params")
	}

	resp, err := s.usecase.SearchMovie(ctx, &entity.SearchMovieRequest{
		Keyword: request.GetKeyword(),
		Page:    int(request.GetPage()),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	movieInfo := []*pb.MovieInfo{}
	for _, v := range resp.Data {
		movieInfo = append(movieInfo, &pb.MovieInfo{
			Title:  v.Title,
			Year:   v.Year,
			ImdbId: v.ImdbID,
			Type:   v.Type,
			Poster: v.Poster,
		})
	}

	return &pb.GetMovieInfoResponse{
		Data: movieInfo,
	}, nil
}
