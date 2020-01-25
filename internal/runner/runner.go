package runner

import (
	"fmt"
	"sync"
	"time"

	"github.com/looplab/fsm"
	log "github.com/sirupsen/logrus"
)

const (
	runEvent     = "run"
	stopEvent    = "stop"
	stoppedEvent = "stopped"
)

// Executable is an object that cun be executed by a Runner
type Executable interface {
	Execute()
}

// Runner is responsible to execute a job periodically
type Runner struct {
	name       string
	executable Executable
	interval   time.Duration
	done       chan bool
	state      *fsm.FSM
	stateMutex sync.Mutex
}

// New returns a new Runner
func New(name string, executable Executable, interval time.Duration) *Runner {
	state := fsm.NewFSM(
		"stopped",
		fsm.Events{
			{Name: runEvent, Src: []string{"stopped"}, Dst: "running"},
			{Name: stopEvent, Src: []string{"running"}, Dst: "stopping"},
			{Name: stoppedEvent, Src: []string{"stopping"}, Dst: "stopped"},
		},
		fsm.Callbacks{},
	)
	return &Runner{
		name:       name,
		executable: executable,
		interval:   interval,
		done:       make(chan bool, 1),
		state:      state,
	}
}

// Run starts the runner. If current state does not allow to start then
// an error will be returned.
// A new go routine will execute the Runnable object every X seconds until Stop() is called.
func (runner *Runner) Run() error {
	runner.stateMutex.Lock()
	defer runner.stateMutex.Unlock()

	if err := runner.state.Event(runEvent); err != nil {
		return fmt.Errorf("can't start runner \"%s\", current state is %s: %s", runner.name, runner.state.Current(), err)
	}

	ticker := time.NewTicker(runner.interval)

	go func() {
		for {
			select {
			case <-ticker.C:
				runner.executable.Execute()
			case <-runner.done:
				ticker.Stop()
				runner.stateMutex.Lock()
				runner.state.Event(stoppedEvent)
				runner.stateMutex.Unlock()

				log.Debugf("runner stopped \"%s\"", runner.name)
				return
			}
		}
	}()

	return nil
}

// Stop stops the runner. If current state does not allow to stop then
// an error will be returned.
func (runner *Runner) Stop() error {
	runner.stateMutex.Lock()
	defer runner.stateMutex.Unlock()
	if err := runner.state.Event(stopEvent); err != nil {
		return fmt.Errorf("can't stop runner \"%s\", current state is %s: %s", runner.name, runner.state.Current(), err)
	}

	log.Debugf("stopping runner \"%s\"", runner.name)
	runner.done <- true
	return nil
}
