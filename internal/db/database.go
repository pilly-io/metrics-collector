package db

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pilly-io/metrics-collector/internal/models"
)

// GormDatabase is wrapper for the orm
type GormDatabase struct {
	*gorm.DB
}

// Database interface
type Database interface {
	Migrate()
	Insert(value interface{})
	DeleteSentMetrics()
	DeleteOldCachedData(interval time.Duration)
}

// New creates an new DB object
func New(DBURI string) (*GormDatabase, error) {
	db, err := gorm.Open("sqlite3", DBURI)
	return &GormDatabase{db}, err
}

// Migrate sync the schemas of the DB
func (db *GormDatabase) Migrate() {
	db.AutoMigrate(&models.Node{}, &models.Namespace{}, &models.Pod{}, &models.PodMetric{})
}

// Insert creates a new record in the right table
func (db *GormDatabase) Insert(value interface{}) {
	db.Create(value)
}

// DeleteSentMetrics remove metrics that were sent from DB
func (db *GormDatabase) DeleteSentMetrics() {
	db.Where("is_sent = true").Delete(&models.PodMetric{})
}

// DeleteOldCachedData removes old data fetched from KubeAPI
func (db *GormDatabase) DeleteOldCachedData(interval time.Duration) {
	maxDate := time.Now().Add(-interval)
	db.Unscoped().Where("created_at < ?", maxDate).Delete(&models.Pod{})
}
