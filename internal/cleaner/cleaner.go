package cleaner

import (
	"github.com/pilly-io/metrics-collector/internal/db"
	"time"
)

type Cleaner struct {
	db              db.Database
	retentionPeriod time.Duration
}

// New creates a new Cleaner
func New(db db.Database, retentionPeriod time.Duration) *Cleaner {
	return &Cleaner{db: db, retentionPeriod: retentionPeriod}
}

// Execute deletes all unecessary data from database
func (cleaner *Cleaner) Execute() {
	cleaner.db.DeleteOldCachedData(cleaner.retentionPeriod)
	cleaner.db.DeleteSentMetrics()
}
