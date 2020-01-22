package collector

import (
	"github.com/pilly-io/metrics-collector/internal/models"
	"github.com/pilly-io/metrics-collector/internal/prometheus/client"
	log "github.com/sirupsen/logrus"
)

// Collector is responsible to fetch metrics periodically and send them back to
// a channel in order for them to be saved.
type Collector struct {
	client  client.Client
	metrics chan<- *models.PodMetric
}

// New returns a new Collector
func New(client client.Client, metrics chan<- *models.PodMetric) *Collector {
	return &Collector{
		client:  client,
		metrics: metrics,
	}
}

// Execute fetches metrics from prometheus
func (collector *Collector) Execute() {
	log.Debug("fetching prometheus metrics")
	collector.fetchMetrics(collector.client.GetPodsCPURequests, "pods CPU requests")
	collector.fetchMetrics(collector.client.GetPodsMemoryRequests, "pods memory requests")
	collector.fetchMetrics(collector.client.GetPodsMemoryUsage, "pods memory usage")
	collector.fetchMetrics(collector.client.GetPodsCPUUsage, "pods CPU usage")
}

func (collector *Collector) fetchMetrics(fetcher func() (client.MetricsList, error), metricName string) {
	if metrics, err := fetcher(); err != nil {
		log.Errorf("failed to fetch %s: %s", metricName, err)
	} else {
		log.Infof("got %d samples for %s", len(metrics), metricName)
		for _, metric := range metrics {
			collector.metrics <- metric
		}
	}
}
