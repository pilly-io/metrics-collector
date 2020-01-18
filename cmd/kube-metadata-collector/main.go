package main

import (
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/pilly-io/metrics-collector/internal/kubernetes"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	//1. Get config
	config := GetConfig()
	log.Info(config)
	//2. Connect to the db
	//db := NewDb(config.DbURI)
	//log.Info(db)
	//3. Initialize Kubernetes API
	kubernetesClient := NewKubernetesClient()
	log.Info(kubernetesClient)
}
