package omdb

import (
	"context"
	"github.com/oky-setiawan/stockbit-test/internal/config"
	"github.com/oky-setiawan/stockbit-test/internal/entity"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"reflect"
	"testing"
)

func Test_omdbGetMovieResponse_getResponse(t *testing.T) {
	type fields struct {
		Data         []omdbGetMovieData
		Total        string
		omdbResponse omdbResponse
	}
	tests := []struct {
		name         string
		fields       fields
		wantResponse *entity.GetMovieResponse
		wantErr      bool
	}{
		{
			name: "Error",
			fields: fields{
				omdbResponse: omdbResponse{
					Response: responseNotOk,
					Error:    "test",
				},
			},
			wantResponse: nil,
			wantErr:      true,
		},
		{
			name: "Success",
			fields: fields{
				omdbResponse: omdbResponse{
					Response: responseOk,
				},
				Data: []omdbGetMovieData{{}},
			},
			wantResponse: &entity.GetMovieResponse{Data: []entity.GetMovieData{{}}},
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &omdbGetMovieResponse{
				Data:         tt.fields.Data,
				Total:        tt.fields.Total,
				omdbResponse: tt.fields.omdbResponse,
			}
			gotResponse, err := o.getResponse()
			if (err != nil) != tt.wantErr {
				t.Errorf("getResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("getResponse() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func Test_omdbDomain_GetMovie(t *testing.T) {
	defer gock.Off()

	t.Run("200 - OK", func(t *testing.T) {
		g := NewGomegaWithT(t)
		o := &omdbDomain{
			HTTPClient: &http.Client{},
			cfg:        &config.OMDBConfig{},
		}

		httpResp := &omdbGetMovieResponse{
			omdbResponse: omdbResponse{
				Response: responseOk,
			},
		}
		gock.New(o.cfg.Host).URL(o.cfg.GetMovieUrl).Reply(http.StatusOK).JSON(httpResp)

		resp, err := o.GetMovie(context.Background(), &entity.GetMovieRequest{})
		g.Expect(err).Should(BeNil())
		g.Expect(resp.Data).Should(HaveLen(len(httpResp.Data)))
	})

	t.Run("200 - NotOK", func(t *testing.T) {
		g := NewGomegaWithT(t)
		o := &omdbDomain{
			HTTPClient: &http.Client{},
			cfg:        &config.OMDBConfig{},
		}

		httpResp := &omdbGetMovieResponse{
			omdbResponse: omdbResponse{
				Response: responseNotOk,
			},
		}
		gock.New(o.cfg.Host).URL(o.cfg.GetMovieUrl).Reply(http.StatusOK).JSON(httpResp)

		resp, err := o.GetMovie(context.Background(), &entity.GetMovieRequest{})
		g.Expect(err).Should(HaveOccurred())
		g.Expect(resp).Should(BeNil())
	})

	t.Run("Error Unmarshal", func(t *testing.T) {
		g := NewGomegaWithT(t)
		o := &omdbDomain{
			HTTPClient: &http.Client{},
			cfg:        &config.OMDBConfig{},
		}

		gock.New(o.cfg.Host).URL(o.cfg.GetMovieUrl).Reply(http.StatusOK).JSON(`}{`)

		resp, err := o.GetMovie(context.Background(), &entity.GetMovieRequest{})
		g.Expect(err).Should(HaveOccurred())
		g.Expect(resp).Should(BeNil())
	})

	t.Run("500", func(t *testing.T) {
		g := NewGomegaWithT(t)
		o := &omdbDomain{
			HTTPClient: &http.Client{},
			cfg:        &config.OMDBConfig{},
		}

		httpResp := &omdbGetMovieResponse{
			omdbResponse: omdbResponse{
				Response: responseNotOk,
			},
		}
		gock.New(o.cfg.Host).URL(o.cfg.GetMovieUrl).Reply(http.StatusInternalServerError).JSON(httpResp)

		resp, err := o.GetMovie(context.Background(), &entity.GetMovieRequest{})
		g.Expect(err).Should(HaveOccurred())
		g.Expect(resp).Should(BeNil())
	})
}
