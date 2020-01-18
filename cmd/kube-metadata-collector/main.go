package main

import (
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/pilly-io/metrics-collector/internal/kubernetes"
	database "github.com/pilly-io/metrics-collector/internal/db"
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
	db, err := database.NewDb(config.DBURI)
	if err != nil {
		log.Fatalf("cannot initialize the database: %s", err)
	}
	db.Migrate()
	//3. Initialize Kubernetes API
	client, err := kubernetes.NewKubernetesClient(config.KubeconfigPath)
	if err != nil {
		log.Fatalf("cannot initialize kubernetes client: %s", err)
	}
	client.ListNodes()
}
