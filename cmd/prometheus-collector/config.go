package main

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	PrometheusURL string
	DBURI         string
	Interval      int
}

func init() {
	viper.SetEnvPrefix("pilly")
	viper.AutomaticEnv()
	viper.SetDefault("INTERVAL", 60)
}

func GetConfig() Config {
	return Config{
		PrometheusURL: viper.GetString("PROMETHEUS_URL"),
		DBURI:         viper.GetString("DB_URI"),
		Interval:      viper.GetInt("INTERVAL"),
	}
}

func (config Config) Validate() error {
	if config.PrometheusURL == "" {
		return errors.New("prometheus URL not set")
	} else if config.DBURI == "" {
		return errors.New("database URI not set")
	}
	return nil
}
