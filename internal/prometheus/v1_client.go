package prometheus

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

const (
	memoryRequestByPodQuery = "sum by (pod, namespace) (container_spec_memory_limit_bytes)"
)

type V1Client struct {
	api v1.API
}

func NewV1Client(endpoint string) (*V1Client, error) {
	apiClient, err := api.NewClient(api.Config{
		Address: endpoint,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating prometheus client: %s", err)
	}

	api := v1.NewAPI(apiClient)
	return &V1Client{api: api}, nil
}

func (client *V1Client) GetPodsMemoryRequests() (error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _, err := client.api.Query(ctx, memoryRequestByPodQuery, time.Now())
	if err != nil {
		return nil, fmt.Errorf("can't fetch pods memory requests: %s", err)
	}
	fmt.Printf("Result:\n%v\n", result)
	return nil, nil
}
