package main

import (
	"errors"

	"github.com/spf13/viper"
)

// Config holds all the configuration values for metrics-controller to run
type Config struct {
	PrometheusURL  string
	KubeconfigPath string
	APIEndpoint    string
	DBURI          string
	Interval       int
}

func init() {
	viper.SetEnvPrefix("pilly")
	viper.AutomaticEnv()
	viper.SetDefault("INTERVAL", 60)
	viper.SetDefault("KUBECONFIG_PATH", nil)
	viper.SetDefault("API_ENDPOINT", "http://127.0.0.1:8080/v1/")
}

// GetConfig fetches the config from  ENV vars and returns a Config
func GetConfig() Config {
	return Config{
		PrometheusURL:  viper.GetString("PROMETHEUS_URL"),
		KubeconfigPath: viper.GetString("KUBECONFIG_PATH"),
		APIEndpoint:    viper.GetString("API_ENDPOINT"),
		DBURI:          viper.GetString("DB_URI"),
		Interval:       viper.GetInt("INTERVAL"),
	}
}

// Validate returns an error if one of the value is invalid
func (config Config) Validate() error {
	if config.PrometheusURL == "" {
		return errors.New("prometheus URL not set")
	} else if config.DBURI == "" {
		return errors.New("database URI not set")
	}
	return nil
}
