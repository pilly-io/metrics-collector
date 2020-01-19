package main

import (
	"os"
	"time"

	"github.com/pilly-io/metrics-collector/internal/prometheus/client"
	"github.com/pilly-io/metrics-collector/internal/prometheus/collector"
	prom "github.com/prometheus/common/model"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	// recup config
	// check que prometheus est up
	// run
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

	metrics := make(chan *prom.Sample, 20)
	collector := collector.New(client, metrics, collector.CollectorConfig{ScrapeInterval: 2 * time.Second})

	if err := collector.Run(); err != nil {
		log.Error(err)
	}
	if err := collector.Run(); err != nil {
		log.Error(err)
	}
	collector.Run()
	go func() {
		for _ = range metrics {
			//log.Infof("%s=%s", sample.Metric, sample.Value)
		}
	}()
	time.Sleep(10 * time.Second)
	if err := collector.Stop(); err != nil {
		log.Error(err)
	}
	if err := collector.Stop(); err != nil {
		log.Error(err)
	}
}
