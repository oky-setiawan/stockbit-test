package core

import (
	"github.com/oky-setiawan/stockbit-test/internal/repository"
	log_repository "github.com/oky-setiawan/stockbit-test/internal/repository/log"
)

type Repositories struct {
	repository.LogRepository
}

func initRepository(driver *Driver) *Repositories {
	logRepo := log_repository.New(&log_repository.Opts{DB: driver.db})

	return &Repositories{
		LogRepository: logRepo,
	}
}
