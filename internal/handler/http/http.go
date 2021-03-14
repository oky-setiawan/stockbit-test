package http

import (
	"github.com/oky-setiawan/stockbit-test/lib/response"
	"net/http"
)

type (
	MovieDelivery interface {
		Get(r *http.Request) *response.JSONResponse
	}
)
