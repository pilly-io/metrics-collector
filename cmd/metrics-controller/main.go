package main

import (
	"os"
	"time"

	"github.com/pilly-io/metrics-collector/internal/db"
	"github.com/pilly-io/metrics-collector/internal/kubernetes"
	"github.com/pilly-io/metrics-collector/internal/models"
	"github.com/pilly-io/metrics-collector/internal/prometheus/client"
	"github.com/pilly-io/metrics-collector/internal/prometheus/collector"
	"github.com/pilly-io/metrics-collector/internal/runner"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func getKubeClient(config Config) kubernetes.IClient {
	var configurator kubernetes.Configurator
	if config.KubeconfigPath != "" {
		configurator = kubernetes.NewBuildFromFlagsConfig(config.KubeconfigPath)
	} else {
		configurator = kubernetes.NewInClusterConfig()
	}
	kubeClient, err := kubernetes.NewKubernetesClient(configurator)
	if err != nil {
		log.Fatalf("can't create kubernetes client: %s", err)
	}

	return kubeClient
}

func getPrometheusRunner(config Config, kubeClient kubernetes.IClient, metrics chan *models.PodMetric) *runner.Runner {
	promConfig := client.ClientConfig{
		Version:  client.APIV1,
		Timeout:  2 * time.Minute,
		Endpoint: config.PrometheusURL,
	}

	client, err := client.New(promConfig)
	if err != nil {
		log.Fatalf("can't create prometheus client: %s", err)
	}

	promCollector := collector.New(client, kubeClient, metrics)
	promRunner := runner.New("prometheus-collector", promCollector, 2*time.Second)

	return promRunner
}

func main() {
	config := GetConfig()
	err := config.Validate()
	if err != nil {
		log.Fatalf("configuration error: %s", err)
	}
	metrics := make(chan *models.PodMetric, 20)

	// Initialize database
	db, err := db.New(config.DBURI)
	if err != nil {
		log.Fatalf("can't create database runner: %s", err)
	}
	db.Migrate()

	kubeClient := getKubeClient(config)

	// Start prometheus collector
	promRunner := getPrometheusRunner(config, kubeClient, metrics)
	if err := promRunner.Run(); err != nil {
		log.Fatalf("can't start prometheus runner: %s", err)
	}

	go func() {
		for metric := range metrics {
			db.Insert(metric)
		}
	}()

	time.Sleep(100 * time.Minute)
	if err := promRunner.Stop(); err != nil {
		log.Error(err)
	}
}
