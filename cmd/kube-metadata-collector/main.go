package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	//1. Get config
	config := GetConfig()
	log.Info(config)
	//2. Connect to the db
	db := NewDb(config.DbURI)
	log.Info(db)
	//3. Initialize Kubernetes API
	kubernetesClient := NewKubernetesClient()
	log.Info(kubernetesClient)
}
