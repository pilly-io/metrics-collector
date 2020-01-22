package main

import (
	database "github.com/pilly-io/metrics-collector/internal/db"
	"github.com/pilly-io/metrics-collector/internal/kubernetes"
	log "github.com/sirupsen/logrus"
	"os"
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
	nodes, err := client.ListNodes()
	if err != nil {
		log.Fatalf("cannot list the nodes: %s", err)
	}
	log.Info(nodes)

	namespaces, err := client.ListNamespaces()
	if err != nil {
		log.Fatalf("cannot list the namespaces: %s", err)
	}
	log.Info(namespaces)

	pods, err := client.ListPods()
	if err != nil {
		log.Fatalf("cannot list the pods: %s", err)
	}
	log.Info(len(*pods))
}
