package client

import (
	"context"
	"fmt"
	"time"

	"github.com/pilly-io/metrics-collector/internal/models"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	prom "github.com/prometheus/common/model"
	log "github.com/sirupsen/logrus"
)

const (
	podsCPURequestsQuery    = "sum by (pod, resource, namespace) (kube_pod_container_resource_requests{resource=\"cpu\"})"
	podsMemoryRequestsQuery = "sum by (pod, resource, namespace) (kube_pod_container_resource_requests{resource=\"memory\"})"
	podsMemoryUsageQuery    = "sum by (pod, namespace) (container_memory_usage_bytes{container!=\"POD\", container=~\".+\"})"
	podsCPUUsageQuery       = "sum (rate(container_cpu_usage_seconds_total{container!=\"POD\", container=~\".+\"}[2m])) by (pod_name, namespace)"
)

type V1Client struct {
	api     v1.API
	timeout time.Duration
}

func NewV1Client(config ClientConfig) (*V1Client, error) {
	apiClient, err := api.NewClient(api.Config{
		Address: config.Endpoint,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating prometheus client: %s", err)
	}

	api := v1.NewAPI(apiClient)
	return &V1Client{api: api, timeout: config.Timeout}, nil
}

func (client *V1Client) GetPodsMemoryRequests() (MetricsList, error) {
	return client.sendQuery(podsMemoryRequestsQuery, "pod_memory_request")
}

func (client *V1Client) GetPodsCPURequests() (MetricsList, error) {
	return client.sendQuery(podsCPURequestsQuery, "pod_cpu_request")
}

func (client *V1Client) GetPodsMemoryUsage() (MetricsList, error) {
	return client.sendQuery(podsMemoryUsageQuery, "pod_memory_usage")
}

func (client *V1Client) GetPodsCPUUsage() (MetricsList, error) {
	return client.sendQuery(podsCPUUsageQuery, "pod_cpu_usage")
}

func (client *V1Client) sendQuery(query string, metricName string) (MetricsList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), client.timeout)
	defer cancel()
	result, warnings, err := client.api.Query(ctx, query, time.Now())
	if err != nil {
		return nil, fmt.Errorf("can't fetch metrics: %s", err)
	}
	if len(warnings) > 0 {
		log.Warningf("got warning while collecting metrics: %s", warnings)
	}

	vector := result.(prom.Vector)
	metrics := make(MetricsList, len(vector))
	for index, sample := range vector {
		metric := models.PodMetric{
			MetricName:  metricName,
			MetricValue: float64(sample.Value),
			Namespace:   string(sample.Metric["namespace"]),
			PodName:     string(sample.Metric["pod"]),
		}
		metrics[index] = &metric
	}
	return metrics, nil
}
