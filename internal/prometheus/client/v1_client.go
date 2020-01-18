package client

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	prom "github.com/prometheus/common/model"
	log "github.com/sirupsen/logrus"
)

const (
	podsRequestsQuery    = "sum by (pod, resource, namespace) (kube_pod_container_resource_requests)"
	podsMemoryUsageQuery = "sum by (pod, namespace) (container_memory_usage_bytes{container!=\"POD\", container=~\".+\"})"
	podsCPUUsageQuery    = "sum (rate(container_cpu_usage_seconds_total{container!=\"POD\", container=~\".+\"}[2m])) by (pod_name, namespace)"
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

func (client *V1Client) GetPodsRequests() (prom.Vector, error) {
	return client.sendQuery(podsRequestsQuery)
}

func (client *V1Client) GetPodsMemoryUsage() (prom.Vector, error) {
	return client.sendQuery(podsMemoryUsageQuery)
}

func (client *V1Client) GetPodsCPUUsage() (prom.Vector, error) {
	return client.sendQuery(podsCPUUsageQuery)
}

func (client *V1Client) sendQuery(query string) (prom.Vector, error) {
	ctx, cancel := context.WithTimeout(context.Background(), client.timeout)
	defer cancel()
	result, warnings, err := client.api.Query(ctx, query, time.Now())
	if err != nil {
		return nil, fmt.Errorf("can't fetch metrics: %s", err)
	}
	if len(warnings) > 0 {
		log.Warningf("got warning while collecting metrics: %s", warnings)
	}
	return result.(prom.Vector), nil
}
