package sender

import (
	"time"

	"github.com/pilly-io/metrics-collector/internal/db"
)

// Config : Sender configuration
type Config struct {
	Endpoint string
	Timeout  time.Duration
}

// Sender : needs the database and config objects
type Sender struct {
	db     *db.GormDatabase
	config *Config
}

//ConsolidateDB : Merge podmetrics records with pods infos
func (sender *Sender) ConsolidateDB() {
	sender.db.Exec(`
		UPDATE pod_metrics SET
			owner_name=(
				SELECT owner_name 
				FROM pods
				WHERE pod_metrics.pod_name=pods.name
			),
			owner_type=(
				SELECT owner_type
				FROM pods
				WHERE pod_metrics.pod_name=pods.name
			)`)
}

//SendAPI : Send the podmetrics to our API
func (sender *Sender) SendAPI() {

}

//NotifyDB : Set IsSent to the podmetrics  records
func (sender *Sender) NotifyDB() {

}

//New : Instanciate a Sender object
func New(db *db.GormDatabase, config *Config) *Sender {
	return &Sender{
		db,
		config,
	}
}
