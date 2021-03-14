package grace

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type HTTPServer struct {
	*http.Server
}

var (
	timeoutHttp = time.Duration(10) * time.Second
)

func ServeHTTP(port string, handler http.Handler) error {

	httpServer := &http.Server{
		Handler:      handler,
		Addr:         port,
		WriteTimeout: time.Duration(5) * time.Second,
		ReadTimeout:  time.Duration(10) * time.Second,
	}
	srv := &HTTPServer{
		Server: httpServer,
	}

	log.Println("http service serve on ", port)
	return srv.ListenAndServe()
}
