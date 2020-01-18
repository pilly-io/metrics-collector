package prometheus

import (
	"fmt"
)

const (
	APIV1 = "v1"
)

type Client interface {
	GetPodsMemoryRequests() (error, error)
}

func New(version string, endpoint string) (Client, error) {
	switch version {
	case APIV1:
		return NewV1Client(endpoint)
	default:
		return nil, fmt.Errorf("%s prometheus API is not supported", version)
	}
}
