package collector

import (
	"github.com/pilly-io/metrics-collector/internal/kubernetes"
	"github.com/pilly-io/metrics-collector/internal/models"
	"github.com/pilly-io/metrics-collector/internal/prometheus/client"
	log "github.com/sirupsen/logrus"
)

// Collector is responsible to fetch metrics periodically and send them back to
// a channel in order for them to be saved.
type Collector struct {
	kubeClient kubernetes.IClient
	client     client.Client
	metrics    chan<- *models.PodMetric
}

// New returns a new Collector
func New(client client.Client, kubeClient kubernetes.IClient, metrics chan<- *models.PodMetric) *Collector {
	return &Collector{
		client:     client,
		metrics:    metrics,
		kubeClient: kubeClient,
	}
}

// Execute fetches metrics from prometheus
func (collector *Collector) Execute() {
	log.Debug("fetching prometheus metrics")
	namespaces, err := collector.kubeClient.ListNamespaces()
	if err != nil {
		log.Errorf("cannot list namespaces to fetch metrics: %s", err)
		return
	}
	for _, namespace := range *namespaces {
		log.Debugf("fetching metrics for namespace=%s", namespace.Name)
		collector.fetchMetrics(collector.client.GetPodsCPURequests, namespace.Name, "pods CPU requests")
		collector.fetchMetrics(collector.client.GetPodsMemoryRequests, namespace.Name, "pods memory requests")
		collector.fetchMetrics(collector.client.GetPodsMemoryUsage, namespace.Name, "pods memory usage")
		collector.fetchMetrics(collector.client.GetPodsCPUUsage, namespace.Name, "pods CPU usage")
	}
}

func (collector *Collector) fetchMetrics(fetcher func(string) (client.MetricsList, error), namespace string, metricName string) {
	if metrics, err := fetcher(namespace); err != nil {
		log.Errorf("failed to fetch %s: %s", metricName, err)
	} else {
		log.Infof("got %d samples for %s", len(metrics), metricName)
		for _, metric := range metrics {
			collector.metrics <- metric
		}
	}
}
