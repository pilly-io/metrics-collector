package client

import (
	"fmt"
	"time"

	prom "github.com/prometheus/common/model"
)

const (
	APIV1 = "v1"
)

type Client interface {
	GetPodsRequests() (prom.Vector, error)
	GetPodsMemoryUsage() (prom.Vector, error)
	GetPodsCPUUsage() (prom.Vector, error)
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
