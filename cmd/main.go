package main

import (
	"github.com/oky-setiawan/stockbit-test/internal/config"
	"github.com/oky-setiawan/stockbit-test/internal/core"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.Get()
	app := core.NewApp(cfg)

	go app.HTTP.Run()

	select {
	case s := <-terminateSignal():
		app.GRPC.CatchSignal(s)
		log.Errorln("Exiting gracefully...", s)
	case err := <-app.HTTP.ListenErr():
		log.Errorln("Error starting http server. exiting gracefully, err: ", err.Error())
	}
}

func terminateSignal() chan os.Signal {
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	return term
}
