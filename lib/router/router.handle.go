package router

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oky-setiawan/stockbit-test/lib/response"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
	"time"
)

type panicObject struct {
	err        interface{}
	stackTrace string
}

type Handle func(r *http.Request) *response.JSONResponse

func (mr *Router) GET(path string, handle Handle) {
	mr.Handle(path, http.MethodGet, handle)
}

func (mr *Router) POST(path string, handle Handle) {
	mr.Handle(path, http.MethodPost, handle)
}

func (mr *Router) PUT(path string, handle Handle) {
	mr.Handle(path, http.MethodPut, handle)
}

func (mr *Router) DELETE(path string, handle Handle) {
	mr.Handle(path, http.MethodDelete, handle)
}

func (mr *Router) Handle(path, method string, handle Handle) {
	fullPath := mr.Options.Prefix + path
	log.Infoln(method, " ", fullPath)
	mr.HttpRouter.Handle(fullPath, mr.handle(handle)).Methods(method)
}

func (mr *Router) handle(handle Handle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*time.Duration(mr.Options.Timeout))

		defer cancel()

		httpParams := mux.Vars(r)

		ctx = context.WithValue(ctx, "HTTPParams", httpParams)

		respChan := make(chan *response.JSONResponse)
		recovered := make(chan panicObject)

		go func() {
			defer func() {
				if err := recover(); err != nil {
					recovered <- panicObject{
						err:        err,
						stackTrace: string(debug.Stack()),
					}
				}
			}()
			respChan <- handle(r)
		}()

		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				response.NewJSONResponse().SetLatency(time.Since(t).Seconds() * 1000).SetError(response.ErrTimeoutError).Send(w)
			}
		case resp := <-respChan:
			resp.SetLatency(time.Since(t).Seconds() * 1000)
			resp.Send(w)
		case cause := <-recovered:
			log.WithFields(log.Fields{
				"path":       r.URL.Path,
				"stackTrace": cause.stackTrace,
				"error":      fmt.Sprintf("%v", cause.err),
			}).Errorln("[Router] panic have occurred")
			response.NewJSONResponse().SetLatency(time.Since(t).Seconds() * 1000).SetError(response.ErrInternalServer).Send(w)
		}

		return
	}
}
