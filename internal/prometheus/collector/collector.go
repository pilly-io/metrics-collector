package collector

import (
	"fmt"
	"sync"
	"time"

	"github.com/looplab/fsm"
	"github.com/pilly-io/metrics-collector/internal/prometheus/client"
	prom "github.com/prometheus/common/model"
	log "github.com/sirupsen/logrus"
)

const (
	runEvent     = "run"
	stopEvent    = "stop"
	stoppedEvent = "stopped"
)

// Config object passed to initialize a new Collector
type Config struct {
	ScrapeInterval time.Duration
}

// Collector is responsible to fetch metrics periodically and send them back to
// a channel in order for them to be saved.
type Collector struct {
	client     client.Client
	metrics    chan<- *prom.Sample
	config     Config
	done       chan bool
	state      *fsm.FSM
	stateMutex sync.Mutex
}

// New returns a new Collector
func New(client client.Client, metrics chan<- *prom.Sample, config Config) *Collector {
	state := fsm.NewFSM(
		"stopped",
		fsm.Events{
			{Name: runEvent, Src: []string{"stopped"}, Dst: "running"},
			{Name: stopEvent, Src: []string{"running"}, Dst: "stopping"},
			{Name: stoppedEvent, Src: []string{"stopping"}, Dst: "stopped"},
		},
		fsm.Callbacks{},
	)
	return &Collector{
		client:  client,
		metrics: metrics,
		config:  config,
		done:    make(chan bool, 1),
		state:   state,
	}
}

// Run starts the collector. If current state does not allow to start then
// an error will be returned.
// A new go routine will start and fetch metrics every X seconds until Stop() is called.
func (collector *Collector) Run() error {
	collector.stateMutex.Lock()
	defer collector.stateMutex.Unlock()

	if err := collector.state.Event(runEvent); err != nil {
		return fmt.Errorf("can't start collector, current state is %s: %s", collector.state.Current(), err)
	}

	ticker := time.NewTicker(collector.config.ScrapeInterval)

	go func() {
		for {
			select {
			case <-ticker.C:
				collector.fetchAllMetrics()
			case <-collector.done:
				ticker.Stop()
				collector.stateMutex.Lock()
				collector.state.Event(stoppedEvent)
				collector.stateMutex.Unlock()

				log.Debug("collector stopped")
				return
			}
		}
	}()

	return nil
}

// Stop stops the collector. If current state does not allow to stop then
// an error will be returned.
func (collector *Collector) Stop() error {
	collector.stateMutex.Lock()
	defer collector.stateMutex.Unlock()
	if err := collector.state.Event(stopEvent); err != nil {
		return fmt.Errorf("can't stop collector, current state is %s: %s", collector.state.Current(), err)
	}

	log.Debug("stopping collector")
	collector.done <- true
	return nil
}

func (collector *Collector) fetchAllMetrics() {
	log.Debug("fetching prometheus metrics")
	collector.fetchMetrics(collector.client.GetPodsRequests, "pods requests")
	collector.fetchMetrics(collector.client.GetPodsMemoryUsage, "pods memory usage")
	collector.fetchMetrics(collector.client.GetPodsCPUUsage, "pods CPU usage")
}

func (collector *Collector) fetchMetrics(fetcher func() (prom.Vector, error), metricName string) {
	if samples, err := fetcher(); err != nil {
		log.Errorf("failed to fetch %s: %s", metricName, err)
	} else {
		log.Infof("got %d samples for %s", len(samples), metricName)
		for _, sample := range samples {
			collector.metrics <- sample
		}
	}
}
