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
	err := config.Validate()
	if err != nil {
		log.Fatalf("configuration error: %s", err)
	}
	//2. Connect to the db
	//db := NewDb(config.DbURI)
	//log.Info(db)
	//3. Initialize Kubernetes API
	kubernetesClient, err := kubernetes.NewKubernetesClient(config.KubeconfigPath)
	if err != nil {
		log.Fatalf("cannot initialize kubernetes client: %s", err)
	}
	log.Info(kubernetesClient)
}
