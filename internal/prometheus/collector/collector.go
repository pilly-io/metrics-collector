package collector

import (
	"github.com/pilly-io/metrics-collector/internal/prometheus/client"
	prom "github.com/prometheus/common/model"
	log "github.com/sirupsen/logrus"
)

// Collector is responsible to fetch metrics periodically and send them back to
// a channel in order for them to be saved.
type Collector struct {
	client  client.Client
	metrics chan<- *prom.Sample
}

// New returns a new Collector
func New(client client.Client, metrics chan<- *prom.Sample) *Collector {
	return &Collector{
		client:  client,
		metrics: metrics,
	}
}

// Execute fetches metrics from prometheus
func (collector *Collector) Execute() {
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
