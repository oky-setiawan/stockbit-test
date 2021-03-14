package config

import "github.com/oky-setiawan/stockbit-test/lib/database"

type Config struct {
	Main    MainConfig    `json:"main"`
	Partner PartnerConfig `json:"partner"`
}

// MainConfig is main config entity
type MainConfig struct {
	Server Server `json:"server"`

	DBConfig database.DBConfig `json:"database"`

	Redis Redis `json:"redis"`
}

// PartnerConfig is partner config entity
type PartnerConfig struct {
	OMDB OMDBConfig `json:"omdb"`
}

type (
	Server struct {
		HTTPPort    string `json:"http_port"`
		HTTPTimeout int    `json:"http_timeout"`
		GRPCPort    string `json:"grpc_port"`
	}

	Redis struct {
		CacheAddress string `json:"cache_address"`
	}
)

type OMDBConfig struct {
	Host        string `json:"host"`
	GetMovieUrl string `json:"get_movie_url"`
	AccessKey   string `json:"access_key"`
}
