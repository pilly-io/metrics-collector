package prometheus

import (
	"fmt"

	prom "github.com/prometheus/common/model"
)

const (
	APIV1 = "v1"
)

type Client interface {
	GetPodsMemoryRequests() (prom.Vector, error)
}

func New(version string, endpoint string) (Client, error) {
	switch version {
	case APIV1:
		return NewV1Client(endpoint)
	default:
		return nil, fmt.Errorf("%s prometheus API is not supported", version)
	}
}
