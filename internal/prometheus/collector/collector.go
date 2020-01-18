package collector

import (
	"time"

	prom "github.com/prometheus/common/model"
	log "github.com/sirupsen/logrus"
	"github.com/pilly-io/metrics-collector/internal/prometheus/client"
)

type CollectorConfig struct {
	ScrapeInterval time.Duration
}

type Collector struct {
	client  client.Client
	metrics chan<- *prom.Sample
	config  CollectorConfig
	done    chan bool
}

func New(client client.Client, metrics chan<- *prom.Sample, config CollectorConfig) *Collector {
	return &Collector{
		client:  client,
		metrics: metrics,
		config:  config,
	}
}

func (collector *Collector) Run() {
	ticker := time.NewTicker(collector.config.ScrapeInterval)

	go func() {
		for {
			select {
			case <-collector.done:
				ticker.Stop()
				log.Debug("collector stopped")
			case <-ticker.C:
				collector.fetchAllMetrics()
			}
		}
	}()
}

func (collector *Collector) Stop() {
	log.Debug("stopping collector")
	collector.done <- true
}

func (collector *Collector) fetchAllMetrics() {
	log.Debug("fetching prometheus metrics")
	collector.fetchMetrics(collector.client.GetPodsRequests, "pods requests")
	collector.fetchMetrics(collector.client.GetPodsMemoryUsage, "pods memory usage")
	collector.fetchMetrics(collector.client.GetPodsCPUUsage, "pods CPU usage")
}

func (collector *Collector) fetchMetrics(fetcher func() (prom.Vector, error), metricName string) {
	if samples, err := fetcher(); err != nil {
		log.Errorf("failed to fetch %s: %s", metricName, err)
	} else {
		log.Infof("got %d samples for %s", len(samples), metricName)
		for _, sample := range samples {
			collector.metrics <- sample
		}
	}
}
