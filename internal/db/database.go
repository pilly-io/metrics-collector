package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pilly-io/metrics-collector/internal/models"
)

// Database is wrapper for the orm
type Database struct {
	*gorm.DB
}

// New creates an new DB object
func New(DBURI string) (*Database, error) {
	db, err := gorm.Open("sqlite3", DBURI)
	return &Database{db}, err
}

// Migrate sync the schemas of the DB
func (db *Database) Migrate() {
	// Dummy migration
	db.AutoMigrate(&models.Node{}, &models.Namespace{}, &models.Pod{}, &models.PodMetric{})
}

// Insert creates a new record in the right table
func (db *Database) Insert(value interface{}) {
	db.Create(value)
}

// DeleteSentMetrics remove metrics that were sent from DB
func (db *Database) DeleteSentMetrics() {
	db.Where("is_sent = true").Delete(&models.PodMetric{})
}

// DeleteOldCachedData removes old data fetched from KubeAPI
func (db *Database) DeleteOldCachedData(interval time.Duration) {
	maxDate := time.Now().Add(-interval)
	fmt.Println(maxDate)
	db.Where("created_at < ?", maxDate).Delete(&models.Pod{})
}
