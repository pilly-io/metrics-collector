package main

import (
	"os"
	"time"

	"github.com/pilly-io/metrics-collector/internal/models"
	"github.com/pilly-io/metrics-collector/internal/prometheus/client"
	"github.com/pilly-io/metrics-collector/internal/prometheus/collector"
	"github.com/pilly-io/metrics-collector/internal/runner"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	config := GetConfig()
	err := config.Validate()
	if err != nil {
		log.Fatalf("configuration error: %s", err)
	}

	promConfig := client.ClientConfig{
		Version:  client.APIV1,
		Timeout:  1 * time.Second,
		Endpoint: config.PrometheusURL,
	}
	client, err := client.New(promConfig)
	if err != nil {
		log.Fatalf("can't create prometheus client: %s", err)
	}

	metrics := make(chan *models.PodMetric, 20)
	promCollector := collector.New(client, metrics)
	promRunner := runner.New("prometheus-collector", promCollector, 2*time.Second)

	if err := promRunner.Run(); err != nil {
		log.Error(err)
	}
	go func() {
		for metric := range metrics {
			log.Infof("%s=%f", metric.MetricName, metric.MetricValue)
		}
	}()
	time.Sleep(10 * time.Second)
	if err := promRunner.Stop(); err != nil {
		log.Error(err)
	}
}
