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
	viper.SetEnvPrefix("pilly")
	viper.AutomaticEnv()
	viper.SetDefault("interval", 60)
	DbURI := viper.GetString("db_uri")
	if DbURI == "" {
		log.Error("DbURI not set")
		os.Exit(1)
	}
	return Config{
		DbURI:    DbURI,
		Interval: viper.GetInt("interval"),
	}
}
