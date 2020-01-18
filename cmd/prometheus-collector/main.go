package main

import (
	"fmt"
	"log"

	prom "github.com/pilly-io/metrics-collector/internal/prometheus"
)

func main() {
	// recup config
	// check que prometheus est up
	// run
	config := GetConfig()
	err := config.Validate()
	if err != nil {
		log.Fatalf("configuration error: %s", err)
	}

	client, err := prom.New(prom.APIV1, config.PrometheusURL)
	if err != nil {
		log.Fatalf("can't create prometheus client: %s", err)
	}

	samples, err := client.GetPodsMemoryRequests()
	for _, sample := range samples {
		fmt.Printf("%s=%s\n", sample.Metric, sample.Value)
	}
}
