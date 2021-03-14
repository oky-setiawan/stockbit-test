package config

import (
	"github.com/oky-setiawan/stockbit-test/lib/files"
	log "github.com/sirupsen/logrus"
)

var (
	configPath = []string{"./files/etc/credentials", "../files/etc/credentials", "../../files/etc/credentials", "C:/Users/Oky Setiawan/go/src/github.com/oky-setiawan/stockbit-test/files/etc/firstapp"}
)

func Get() *Config {
	cfg := &Config{}

	//read main
	err := files.ReadJSON("app.json", configPath, &cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"error":      err.Error(),
			"configPath": configPath,
		}).Fatalf("failed read config file")
	} else {
		log.Infoln("success read config file")
	}

	return cfg
}
