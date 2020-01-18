package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DbURI    string
	Interval int
}

func GetConfig() Config {
	return Config{
		DBURI:         viper.GetString("DB_URI"),
		Interval:      viper.GetInt("INTERVAL"),
	}
}

func (config Config) Validate() error {
	if if config.DBURI == "" {
		return errors.New("database URI not set")
	}
	return nil
}