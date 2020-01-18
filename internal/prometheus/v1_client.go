package prometheus

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
	memoryRequestByPodQuery = "sum by (pod, resource, namespace) (kube_pod_container_resource_requests)"
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

func (client *V1Client) GetPodsMemoryRequests() (prom.Vector, error) {
	return client.sendQuery(memoryRequestByPodQuery)
}

func (client *V1Client) sendQuery(query string) (prom.Vector, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
