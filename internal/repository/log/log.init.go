package log

import "github.com/oky-setiawan/stockbit-test/lib/database"

type Opts struct {
	DB *database.Store
}

type logRepository struct {
	db *database.Store
}

func New(o *Opts) *logRepository {
	return &logRepository{
		db: o.DB,
	}
}
