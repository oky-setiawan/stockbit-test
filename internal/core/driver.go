package core

import (
	"github.com/oky-setiawan/stockbit-test/internal/config"
	"github.com/oky-setiawan/stockbit-test/lib/database"
)

type Driver struct {
	db *database.Store
}

func initDriver(cfg *config.Config) *Driver {

	// Database
	dataStore := database.New(cfg.Main.DBConfig, database.DriverMySQL)

	return &Driver{
		db: dataStore,
	}
}
