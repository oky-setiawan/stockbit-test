package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	HttpRouter     *mux.Router
	WrappedHandler http.HandlerFunc
	Options        *Options
}

type Options struct {
	Prefix  string
	Timeout int
}

var (
	httpRouter *mux.Router
)

func init() {
	httpRouter = mux.NewRouter()
}

func New(o *Options) *Router {
	router := &Router{
		HttpRouter: httpRouter,
		Options:    o,
	}
	return router
}

func GetHandler() *mux.Router {
	return httpRouter
}
