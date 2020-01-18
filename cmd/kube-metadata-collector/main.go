package main

import (
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/pilly-io/metrics-collector/internal/kubernetes"
	"github.com/spf13/viper"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	viper.SetEnvPrefix("pilly")
	viper.AutomaticEnv()
	viper.SetDefault("interval", 60)
}

func main() {
	//1. Get config
	config := GetConfig()
	err := config.Validate()
	if err != nil {
		log.Fatalf("configuration error: %s", err)
	}
	//2. Connect to the db
	//db := NewDb(config.DbURI)
	//log.Info(db)
	//3. Initialize Kubernetes API
	kubernetesClient := NewKubernetesClient()
	log.Info(kubernetesClient)
}
