package main

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	DBURI    string
	KubeconfigPath string
	Interval int
}

func init() {
	viper.SetEnvPrefix("pilly")
	viper.AutomaticEnv()
	viper.SetDefault("kubeconfig_path", nil)
	viper.SetDefault("interval", 60)
}

func GetConfig() Config {
	return Config{
		DBURI:         viper.GetString("DB_URI"),
		KubeconfigPath: viper.GetString("KUBECONFIG_PATH"),
		Interval:      viper.GetInt("INTERVAL"),
	}
}

func (config Config) Validate() error {
	if config.DBURI == "" {
		return errors.New("database URI not set")
	}
	return nil
}