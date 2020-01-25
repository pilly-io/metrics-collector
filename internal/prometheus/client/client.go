package client

import (
	"fmt"
	"time"

	"github.com/pilly-io/metrics-collector/internal/models"
)

const (
	APIV1 = "v1"
)

type MetricsList []*models.PodMetric

type Client interface {
	GetPodsCPURequests(namespace string) (MetricsList, error)
	GetPodsMemoryRequests(namespace string) (MetricsList, error)
	GetPodsMemoryUsage(namespace string) (MetricsList, error)
	GetPodsCPUUsage(namespace string) (MetricsList, error)
}

type ClientConfig struct {
	Endpoint string
	Version  string
	Timeout  time.Duration
}

func New(config ClientConfig) (Client, error) {
	switch config.Version {
	case APIV1:
		return NewV1Client(config)
	default:
		return nil, fmt.Errorf("%s prometheus API is not supported", config.Version)
	}
}
